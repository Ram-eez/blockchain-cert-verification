package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthorizeJWT(c *gin.Context)
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (m *middleware) AuthorizeJWT(
	c *gin.Context,
) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "missing authorization header",
			},
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
