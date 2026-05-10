package db

import (
	"blockchain/internal/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, cnf *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cnf.PgUsername,
		cnf.PgPassword,
		cnf.PgHost,
		cnf.PgPort,
		cnf.PgDatabaseName,
		cnf.PgSSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("failed to parse postgres config: ", err)
	}

	// now configure PGX Pool

	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1

	poolConfig.MaxConnIdleTime = 5 * time.Minute
	poolConfig.MaxConnLifetime = 30 * time.Minute

	poolConfig.HealthCheckPeriod = 1 * time.Minute

	// create a new pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal("failed to create postgres pool: ", err)
	}

	// now verify pool conn

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	log.Println("connected to postgres successfully ")

	return pool
}
