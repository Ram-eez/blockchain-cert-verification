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
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "missing token"},
		)
		return
	}

	instituteID, err := m.ValidateJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid token"},
		)
		return
	}

	c.Set("institute_id", instituteID)

	c.Next()
}

func (m *middleware) ValidateJWT(tokenString string) (uuid.UUID, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cnf.Secret), nil
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	instituteID, err := uuid.Parse(claims["institute_id"].(string))
	if err != nil {
		return uuid.Nil, err
	}

	return instituteID, nil
}

func (m *middleware) GenerateJWT(instituteID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"institute_id": instituteID.String(),
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(m.cnf.Secret))
}
