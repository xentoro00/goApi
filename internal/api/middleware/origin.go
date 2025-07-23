package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckOrigin creates a middleware that checks the request's Origin header.
func CheckOrigin() gin.HandlerFunc {
	// Define the list of allowed frontend origins.
	allowedOrigins := []string{"http://localhost:3000"} // Add your production domain here later

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Allow requests with no origin (e.g., server-to-server, mobile apps)
		// or requests from allowed origins.
		isAllowed := false
		for _, o := range allowedOrigins {
			if o == origin {
				isAllowed = true
				break
			}
		}

		// Browsers always send an Origin header for cross-origin requests.
		// If the origin is present but not allowed, block it.
		if origin != "" && !isAllowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied: invalid origin"})
			return
		}

		c.Next()
	}
}
