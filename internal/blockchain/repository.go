package blockchain

import "github.com/jackc/pgx/v5/pgxpool"

type ProjectRepository interface {
}

type projectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}
