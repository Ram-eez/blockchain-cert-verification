package bootstrap

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/config"
	"blockchain/internal/db"
	"blockchain/internal/handlers"
	"blockchain/internal/middleware"
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

		// Initialize JWT
	jwt := middleware.NewMiddleware(cfg)

	// Initialize Blockchain Service
	blockchainService :=
		blockchain.NewBlockChainService(cfg, projectRepository, jwt)

	// Initialize Handlers
	projectHandlers := handlers.NewProjectHandler(blockchainService, jwt, *cfg)

	// Initialize Router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Mount Routes
	projectHandlers.MountRoutes(router)

	// Create HTTP Server
	httpServer := server.NewServer(cfg, router, pool)

	return httpServer, nil
}
