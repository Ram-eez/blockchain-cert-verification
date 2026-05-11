package handlers

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/middleware"
	"crypto/sha256"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type projectHandlers struct {
	services   blockchain.BlockChainService
	middleware middleware.Middleware
}

func NewProjectHandler(services blockchain.BlockChainService, middleware middleware.Middleware) *projectHandlers {
	return &projectHandlers{
		services:   services,
		middleware: middleware,
	}
}

func (pH *projectHandlers) MountRoutes(router *gin.Engine) {
	public := router.Group("/")
	protected := router.Group("/")

	protected.Use(pH.middleware.AuthorizeJWT)

	public.POST("/certificates/verify")
	protected.POST("/certificates/issue")
	protected.POST("/certificates/rewoke")

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func (pH *projectHandlers) IssueCertificate(c *gin.Context) {
	pdfFile, err := c.FormFile("pdf")
	if err != nil {
		c.String(http.StatusBadRequest, "pdf file is required")
		return
	}

	file, err := pdfFile.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to open pdf file")
		return
	}

	defer file.Close()

	// hash certificate
	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to hash pdf")
		return
	}

	var pdfHash [32]byte
	copy(pdfHash[:], hasher.Sum(nil))
	req := blockchain.IssueCertificateRequest{
		PdfHash:          pdfHash,
		RecipientName:    c.PostForm("recipient_name"),
		CourseName:       c.PostForm("course_name"),
		Grade:            c.PostForm("grade"),
		IssuingAuthority: c.PostForm("issuing_authority"),
	}

	err = pH.services.IssueCertificate(c.Request.Context(), req)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.String(
		http.StatusOK,
		"certificate issued successfully",
	)
}
