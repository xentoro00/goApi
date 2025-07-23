package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckAPIKey creates a middleware that validates the X-API-Key header.
func CheckAPIKey(expectedKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		if key != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			return
		}

		c.Next()
	}
}
