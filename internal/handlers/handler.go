package handlers

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/middleware"

	"github.com/gin-gonic/gin"
)

type projectHandlers struct {
	services   blockchain.BlockChainService
	middleware middleware.Middleware
}

func NewProjectHandler(services blockchain.BlockChainService) *projectHandlers {
	return &projectHandlers{
		services:   services,
		middleware: middleware.Middleware{},
	}
}

func (pH *projectHandlers) MountRoutes(router *gin.Engine) {
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
