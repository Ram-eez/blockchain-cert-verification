package migrator

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const migrationsTable = "schema_migrations"

type Direction string

const (
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

type Migration struct {
	Version   string
	Name      string
	Direction Direction
	Path      string
}

type AppliedMigration struct {
	Version   string
	Name      string
	AppliedAt string
}

func Up(ctx context.Context, pool *pgxpool.Pool, dir string) error {
	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return err
	}

	migrations, err := LoadMigrations(dir, DirectionUp)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		applied, err := isApplied(ctx, pool, migration.Version)
		if err != nil {
			return err
		}
		if applied {
			continue
		}

		if err := applyUpMigration(ctx, pool, migration); err != nil {
			return err
		}
	}

	return nil
}

func Down(ctx context.Context, pool *pgxpool.Pool, dir string, steps int) error {
	if steps < 1 {
		return fmt.Errorf("steps must be at least 1")
	}

	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return err
	}

	applied, err := appliedMigrations(ctx, pool)
	if err != nil {
		return err
	}

	if len(applied) == 0 {
		return nil
	}

	downMigrations, err := LoadMigrations(dir, DirectionDown)
	if err != nil {
		return err
	}

	downByVersion := make(map[string]Migration, len(downMigrations))
	for _, migration := range downMigrations {
		downByVersion[migration.Version] = migration
	}

	if steps > len(applied) {
		steps = len(applied)
	}

	for i := 0; i < steps; i++ {
		current := applied[len(applied)-1-i]
		migration, ok := downByVersion[current.Version]
		if !ok {
			return fmt.Errorf("missing down migration for version %s", current.Version)
		}

		if err := applyDownMigration(ctx, pool, migration); err != nil {
			return err
		}
	}

	return nil
}

func Status(ctx context.Context, pool *pgxpool.Pool, dir string) ([]string, error) {
	if err := ensureMigrationsTable(ctx, pool); err != nil {
		return nil, err
	}

	migrations, err := LoadMigrations(dir, DirectionUp)
	if err != nil {
		return nil, err
	}

	appliedSet := make(map[string]AppliedMigration)
	applied, err := appliedMigrations(ctx, pool)
	if err != nil {
		return nil, err
	}
	for _, migration := range applied {
		appliedSet[migration.Version] = migration
	}

	lines := make([]string, 0, len(migrations))
	for _, migration := range migrations {
		if existing, ok := appliedSet[migration.Version]; ok {
			lines = append(lines, fmt.Sprintf("[applied] %s %s (%s)", migration.Version, migration.Name, existing.AppliedAt))
			continue
		}

		lines = append(lines, fmt.Sprintf("[pending] %s %s", migration.Version, migration.Name))
	}

	return lines, nil
}

func LoadMigrations(dir string, direction Direction) ([]Migration, error) {
	entries := make([]Migration, 0)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}

		name := d.Name()
		suffix := "." + string(direction) + ".sql"
		if !strings.HasSuffix(name, suffix) {
			return nil
		}

		migration, err := parseMigrationName(path, name, direction)
		if err != nil {
			return err
		}

		entries = append(entries, migration)
		return nil
	})
	if err != nil {
		return nil, err
	}

	slices.SortFunc(entries, func(a, b Migration) int {
		return strings.Compare(a.Version, b.Version)
	})

	return entries, nil
}

func parseMigrationName(path, name string, direction Direction) (Migration, error) {
	suffix := "." + string(direction) + ".sql"
	base := strings.TrimSuffix(name, suffix)
	parts := strings.SplitN(base, "_", 2)
	if len(parts) != 2 {
		return Migration{}, fmt.Errorf("invalid migration file name %q", name)
	}

	return Migration{
		Version:   parts[0],
		Name:      parts[1],
		Direction: direction,
		Path:      path,
	}, nil
}

func ensureMigrationsTable(ctx context.Context, pool *pgxpool.Pool) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
		)
	`, migrationsTable)

	_, err := pool.Exec(ctx, query)
	return err
}

func isApplied(ctx context.Context, pool *pgxpool.Pool, version string) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE version = $1)`, migrationsTable)

	var exists bool
	if err := pool.QueryRow(ctx, query, version).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func appliedMigrations(ctx context.Context, pool *pgxpool.Pool) ([]AppliedMigration, error) {
	query := fmt.Sprintf(`
		SELECT version, name, to_char(applied_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')
		FROM %s
		ORDER BY version ASC
	`, migrationsTable)

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	migrations := make([]AppliedMigration, 0)
	for rows.Next() {
		var migration AppliedMigration
		if err := rows.Scan(&migration.Version, &migration.Name, &migration.AppliedAt); err != nil {
			return nil, err
		}
		migrations = append(migrations, migration)
	}

	return migrations, rows.Err()
}

func applyUpMigration(ctx context.Context, pool *pgxpool.Pool, migration Migration) error {
	contents, err := os.ReadFile(migration.Path)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", migration.Path, err)
	}

	return withTx(ctx, pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, string(contents)); err != nil {
			return fmt.Errorf("execute migration %s: %w", migration.Path, err)
		}

		query := fmt.Sprintf(`INSERT INTO %s (version, name) VALUES ($1, $2)`, migrationsTable)
		if _, err := tx.Exec(ctx, query, migration.Version, migration.Name); err != nil {
			return fmt.Errorf("record migration %s: %w", migration.Path, err)
		}

		return nil
	})
}

func applyDownMigration(ctx context.Context, pool *pgxpool.Pool, migration Migration) error {
	contents, err := os.ReadFile(migration.Path)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", migration.Path, err)
	}

	return withTx(ctx, pool, func(tx pgx.Tx) error {
		sql := strings.TrimSpace(string(contents))
		if sql != "" {
			if _, err := tx.Exec(ctx, sql); err != nil {
				return fmt.Errorf("execute migration %s: %w", migration.Path, err)
			}
		}

		query := fmt.Sprintf(`DELETE FROM %s WHERE version = $1`, migrationsTable)
		if _, err := tx.Exec(ctx, query, migration.Version); err != nil {
			return fmt.Errorf("delete migration record %s: %w", migration.Path, err)
		}

		return nil
	})
}

func withTx(ctx context.Context, pool *pgxpool.Pool, fn func(tx pgx.Tx) error) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := fn(tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		if errors.Is(err, pgx.ErrTxClosed) {
			return nil
		}
		return err
	}

	return nil
}
