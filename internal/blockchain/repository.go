package blockchain

import "github.com/jackc/pgx/v5/pgxpool"

type BlockChainRepository interface {
}

type blockChainRepository struct {
	db *pgxpool.Pool
}

func NewBlockChainRepository(db *pgxpool.Pool) BlockChainRepository {
	return &blockChainRepository{
		db: db,
	}
}
