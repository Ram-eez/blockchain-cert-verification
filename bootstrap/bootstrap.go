package bootstrap

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/config"
	"blockchain/internal/db"
	"blockchain/internal/handlers"
	"blockchain/internal/server"
	"context"

	"github.com/gin-gonic/gin"
)

func InitializeApplication(cfg *config.Config) (*server.Server, error) {

	// Root Context
	ctx := context.Background()

	// Initialize PostgreSQL
	pool := db.NewPostgresPool(ctx, cfg)

	// Initialize Repositories
	projectRepository :=
		blockchain.NewProjectRepository(pool)

	// Initialize Blockchain Service
	blockchainService :=
		blockchain.NewBlockChainService(cfg, projectRepository)

	// Initialize Handlers
	projectHandlers := handlers.NewProjectHandler(blockchainService)

	// Initialize Router
	router := gin.Default()

	// Mount Routes
	projectHandlers.MountRoutes(router)

	// Create HTTP Server
	httpServer := server.NewServer(cfg, router, pool)

	return httpServer, nil
}
