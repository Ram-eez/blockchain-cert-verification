package middleware

import (
	"blockchain/internal/config"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Middleware interface {
	// Institute
	AuthorizeJWT(c *gin.Context)
	GenerateJWT(instituteID uuid.UUID, instituteName string) (string, error)
	GenerateAdminJWT() (string, error)
	// ADMIN
	AuthorizeAdmin(c *gin.Context)
	ValidateJWT(tokenString string) (uuid.UUID, string, error)
	ValidateAdminJWT(tokenString string) error
}

type middleware struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) Middleware {
	return &middleware{
		cnf: cnf,
	}
}

func (m *middleware) AuthorizeJWT(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "missing token"},
		)
		return
	}

	instituteID, instituteName, err := m.ValidateJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid token"},
		)
		return
	}

	c.Set("institute_id", instituteID)
	c.Set("institute_name", instituteName)
	c.Next()
}

func (m *middleware) ValidateJWT(tokenString string) (uuid.UUID, string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cnf.Secret), nil
		},
	)

	if err != nil {
		return uuid.Nil, "", err
	}

	if !token.Valid {
		return uuid.Nil, "", errors.New("invalid token")
	}

	instituteID, err := uuid.Parse(claims["institute_id"].(string))
	if err != nil {
		return uuid.Nil, "", err
	}

	instituteName, ok := claims["institute_name"].(string)
	if !ok {
		return uuid.Nil, "", errors.New("missing institute name")
	}

	return instituteID, instituteName, nil
}

func (m *middleware) GenerateJWT(instituteID uuid.UUID, instituteName string) (string, error) {
	claims := jwt.MapClaims{
		"institute_id":   instituteID.String(),
		"institute_name": instituteName,
		"exp":            time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(m.cnf.Secret))
}

func (m *middleware) GenerateAdminJWT() (string, error) {
	claims := jwt.MapClaims{
		"is_admin": true,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.cnf.Secret))
}

func (m *middleware) ValidateAdminJWT(tokenString string) error {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.cnf.Secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	isAdmin, ok := claims["is_admin"].(bool)
	if !ok || !isAdmin {
		return errors.New("not admin")
	}

	return nil
}

func (m *middleware) AuthorizeAdmin(c *gin.Context) {
	token, err := c.Cookie("admin_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing admin token"})
		return
	}

	err = m.ValidateAdminJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid admin token"})
		return
	}

	c.Next()
}
