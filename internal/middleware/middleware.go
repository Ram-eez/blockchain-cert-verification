package middleware

import (
	"blockchain/internal/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Middleware interface {
	AuthorizeJWT(c *gin.Context)
	GenerateJWT(instituteID uuid.UUID) (string, error)
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

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "missing authorization header"},
		)
		return
	}

	token := strings.TrimPrefix(
		authHeader,
		"Bearer ",
	)

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "invalid token",
			},
		)
		return
	}

	// TODO:
	// validate JWT properly later

	c.Set("institute_id", "temp-institute-id")

	c.Next()
}

func (m *middleware) GenerateJWT(instituteID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"institute_id": instituteID.String(),
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.cnf.Secret))
}
