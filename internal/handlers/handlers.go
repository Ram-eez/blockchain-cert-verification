package handlers

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/middleware"
	"crypto/sha256"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// auth
	public.GET("/login", pH.LoginPage)
	public.POST("/login", pH.Login)
	//public.GET("/logout", pH.Logout)

	// pages
	public.GET("/", pH.VerifyPage)
	public.GET("/verify", pH.VerifyPage)

	protected.GET("/issue", pH.IssuePage)
	protected.GET("/revoke", pH.RevokePage)

	// certificate apis
	public.GET("/certificates/verify/:hash", pH.VerifyCertificateByURL)
	public.POST("/certificates/verify/hash", pH.VerifyCertificateByHash)
	public.POST("/certificates/verify", pH.VerifyCertificate)
	protected.POST("/certificates/issue", pH.IssueCertificate)
	protected.POST("/certificates/revoke", pH.RevokeCertificate)
	protected.POST("/certificates/revoke/hash", pH.RevokeCertificateByHash)

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

	instituteID, ok := c.Get("institute_id")
	if !ok {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	instituteName, ok := c.Get("institute_name")
	if !ok {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	req := blockchain.IssueCertificateRequest{
		InstituteID:      instituteID.(uuid.UUID),
		IssuingAuthority: instituteName.(string),
		PdfHash:          pdfHash,
		RecipientName:    c.PostForm("recipient_name"),
		CourseName:       c.PostForm("course_name"),
		Grade:            c.PostForm("grade"),
	}

	resp, err := pH.services.IssueCertificate(c.Request.Context(), req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.HTML(http.StatusOK, "issue_result.html", resp)
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

	pH.verifyHash(c, pdfHash)
}

func (pH *projectHandlers) verifyHash(c *gin.Context, pdfHash [32]byte) {
	result, err := pH.services.VerifyCertificate(
		c.Request.Context(),
		pdfHash,
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

func (pH *projectHandlers) VerifyCertificateByHash(c *gin.Context) {
	hashHex := c.PostForm("certificate_hash")

	hash := common.HexToHash(hashHex)

	var pdfHash [32]byte
	copy(pdfHash[:], hash.Bytes())

	pH.verifyHash(c, pdfHash)
}

func (pH *projectHandlers) VerifyCertificateByURL(c *gin.Context) {
	hashHex := c.Param("hash")

	hash := common.HexToHash(hashHex)

	var pdfHash [32]byte
	copy(pdfHash[:], hash.Bytes())

	pH.verifyHash(c, pdfHash)
}

func (pH *projectHandlers) revokeHash(c *gin.Context, pdfHash [32]byte) {
	err := pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
	)

	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.String(
		http.StatusOK,
		"certificate revoked successfully",
	)

}

func (pH *projectHandlers) RevokeCertificate(c *gin.Context) {
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

	err = pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
	)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.String(
		http.StatusOK,
		"certificate revoked successfully",
	)
}

func (pH *projectHandlers) RevokeCertificateByHash(c *gin.Context) {
	hashHex := c.PostForm("certificate_hash")

	if hashHex == "" {
		c.String(
			http.StatusBadRequest,
			"certificate hash is required",
		)
		return
	}

	hash := common.HexToHash(hashHex)

	var pdfHash [32]byte
	copy(pdfHash[:], hash.Bytes())

	err := pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
	)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.String(
		http.StatusOK,
		"certificate revoked successfully",
	)
}

// HTMX Pages and Endpoints
func (pH *projectHandlers) IssuePage(c *gin.Context) {
	c.HTML(http.StatusOK, "issue.html", nil)
}

func (pH *projectHandlers) VerifyPage(c *gin.Context) {
	c.HTML(http.StatusOK, "verify.html", nil)
}

func (pH *projectHandlers) RevokePage(c *gin.Context) {
	c.HTML(http.StatusOK, "revoke.html", nil)
}
func (pH *projectHandlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (pH *projectHandlers) Login(c *gin.Context) {
	req := blockchain.LoginRequest{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	token, err := pH.services.Login(c.Request.Context(), req)
	if err != nil {
		c.String(http.StatusUnauthorized, "invalid credentials")
		return
	}

	c.SetCookie(
		"token",
		token,
		86400,
		"/",
		"",
		false,
		true,
	)

	c.Header("HX-Redirect", "/issue")
	c.String(http.StatusOK, "")
}
