package handlers

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/config"
	"blockchain/internal/middleware"
	"crypto/sha256"
	"errors"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type projectHandlers struct {
	services   blockchain.BlockChainService
	middleware middleware.Middleware
	cnf        config.Config
	jwt        middleware.Middleware
}

func NewProjectHandler(services blockchain.BlockChainService, middleware middleware.Middleware, cnf config.Config) *projectHandlers {
	return &projectHandlers{
		services:   services,
		middleware: middleware,
		cnf:        cnf,
		jwt:        middleware,
	}
}
func (pH *projectHandlers) MountRoutes(router *gin.Engine) {
	public := router.Group("/")
	protected := router.Group("/")
	admin := router.Group("/admin")

	protected.Use(pH.middleware.AuthorizeJWT)

	admin.Use(pH.middleware.AuthorizeAdmin)

	// auth
	public.GET("/", pH.HomePage)
	public.GET("/login", pH.LoginPage)
	public.POST("/login", pH.Login)
	//public.GET("/logout", pH.Logout)

	// pages
	public.GET("/verify", pH.VerifyPage)
	public.GET("/verify/hash", pH.VerifyCertificateByHashPage)
	public.GET("/verify/:hash", pH.VerifyCertificateByURL)

	protected.GET("/issue", pH.IssuePage)
	protected.GET("/revoke", pH.RevokePage)

	// certificate apis
	public.POST("/certificates/verify", pH.VerifyCertificate)

	protected.POST("/certificates/issue", pH.IssueCertificate)
	protected.POST("/certificates/revoke", pH.RevokeCertificate)
	protected.POST("/certificates/revoke/hash", pH.RevokeCertificateByHash)

	// admin auth
	public.GET("/admin/login", pH.AdminLoginPage)
	public.POST("/admin/login", pH.AdminLogin)

	// admin pages
	admin.GET("/dashboard", pH.AdminDashboard)

	// admin apis
	admin.POST("/institutions", pH.CreateInstitution)

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func (h *projectHandlers) AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != h.cnf.AdminUsername || password != h.cnf.AdminPassword {
		c.String(http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := h.jwt.GenerateAdminJWT()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie("admin_token", token, 86400, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

func (pH *projectHandlers) CreateInstitution(c *gin.Context) {
	var req blockchain.CreateInstitutionRequest

	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := pH.services.CreateInstitution(c.Request.Context(), req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "institution_result.html", resp)
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

	c.Redirect(http.StatusSeeOther, "/verify/"+common.Bytes2Hex(pdfHash[:]))
}

func (pH *projectHandlers) verifyHash(c *gin.Context, pdfHash [32]byte) {
	result, err := pH.services.VerifyCertificate(
		c.Request.Context(),
		pdfHash,
	)

	if err != nil {
		c.HTML(http.StatusNotFound, "verify_result.html", &blockchain.VerifyCertificateResponse{Exists: false})
		return
	}

	c.HTML(http.StatusOK, "verify_result.html", result)
}

func (pH *projectHandlers) VerifyCertificateByHashPage(c *gin.Context) {
	hashHex := c.Query("certificate_hash")

	if hashHex == "" {
		c.Redirect(http.StatusSeeOther, "/verify")
		return
	}

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

func (pH *projectHandlers) revokeHash(c *gin.Context, pdfHash [32]byte, instituteID uuid.UUID) {
	err := pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
		instituteID,
	)

	if err != nil {
		switch {
		case errors.Is(err, blockchain.ErrCertificateAlreadyRevoked):
			c.String(http.StatusConflict, err.Error())
		case errors.Is(err, blockchain.ErrCertificateNotFound):
			c.String(http.StatusNotFound, err.Error())
		case errors.Is(err, blockchain.ErrCertificateUnauthorized):
			c.String(http.StatusForbidden, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
		}
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

	instituteID, ok := c.Get("institute_id")
	if !ok {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	err = pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
		instituteID.(uuid.UUID),
	)
	if err != nil {
		switch {
		case errors.Is(err, blockchain.ErrCertificateAlreadyRevoked):
			c.String(http.StatusConflict, err.Error())
		case errors.Is(err, blockchain.ErrCertificateNotFound):
			c.String(http.StatusNotFound, err.Error())
		case errors.Is(err, blockchain.ErrCertificateUnauthorized):
			c.String(http.StatusForbidden, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
		}
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

	instituteID, ok := c.Get("institute_id")
	if !ok {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	err := pH.services.RevokeCertificate(
		c.Request.Context(),
		pdfHash,
		instituteID.(uuid.UUID),
	)
	if err != nil {
		switch {
		case errors.Is(err, blockchain.ErrCertificateAlreadyRevoked):
			c.String(http.StatusConflict, err.Error())
		case errors.Is(err, blockchain.ErrCertificateNotFound):
			c.String(http.StatusNotFound, err.Error())
		case errors.Is(err, blockchain.ErrCertificateUnauthorized):
			c.String(http.StatusForbidden, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
		}
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
func (pH *projectHandlers) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func (pH *projectHandlers) VerifyPage(c *gin.Context) {
	c.HTML(http.StatusOK, "verify.html", gin.H{
		"can_manage": pH.hasInstitutionSession(c),
	})
}
func (pH *projectHandlers) RevokePage(c *gin.Context) {
	c.HTML(http.StatusOK, "revoke.html", nil)
}
func (pH *projectHandlers) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func (pH *projectHandlers) AdminLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}
func (pH *projectHandlers) AdminDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_dashboard.html", nil)
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

func (pH *projectHandlers) hasInstitutionSession(c *gin.Context) bool {
	token, err := c.Cookie("token")
	if err != nil {
		return false
	}

	_, _, err = pH.middleware.ValidateJWT(token)
	return err == nil
}
