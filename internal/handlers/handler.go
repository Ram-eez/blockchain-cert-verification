package handlers

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/middleware"
	"crypto/sha256"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
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

	public.GET("/certificates/verify/:hash", pH.VerifyCertificateByHash)
	public.POST("/certificates/verify", pH.VerifyCertificate)
	protected.POST("/certificates/issue", pH.IssueCertificate)
	protected.POST("/certificates/rewoke", pH.RevokeCertificate)

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

	qrCode, err := pH.services.IssueCertificate(c.Request.Context(), req)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.Data(http.StatusOK, "image/png", qrCode.QRCode)
}

func (pH *projectHandlers) VerifyCertificate(c *gin.Context) {
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

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to hash pdf")
		return
	}

	var pdfHash [32]byte
	copy(pdfHash[:], hasher.Sum(nil))
	result, err := pH.services.VerifyCertificate(c.Request.Context(), pdfHash)
	if err != nil {
		c.String(http.StatusNotFound, "certificate not found")
		return
	}

	if !result.IsValid {
		c.String(http.StatusOK, "certificate is revoked")
		return
	}

	c.String(http.StatusOK, "certificate is valid")
}

func (pH *projectHandlers) VerifyCertificateByHash(c *gin.Context) {
	hashHex := c.Param("hash")
	hash := common.HexToHash(hashHex)
	result, err := pH.services.VerifyCertificate(
		c.Request.Context(),
		hash,
	)

	if err != nil {
		c.String(http.StatusNotFound, "certificate not found")
		return
	}

	if !result.IsValid {
		c.String(http.StatusOK, "certificate revoked")
		return
	}

	c.String(http.StatusOK, "certificate valid")
}

func (pH *projectHandlers) RevokeCertificate(c *gin.Context) {
	hashHex := c.PostForm("certificate_hash")

	if hashHex == "" {
		c.String(http.StatusBadRequest, "certificate hash is required")
		return
	}

	certificateHash := common.HexToHash(hashHex)

	err := pH.services.RevokeCertificate(c.Request.Context(), certificateHash)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "certificate revoked successfully")
}
