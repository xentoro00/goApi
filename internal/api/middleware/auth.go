package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	supaauth "github.com/supabase-community/auth-go"
)

func Protected(supaClient supaauth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no access token found"})
			return
		}

		authedClient := supaClient.WithToken(token)
		user, err := authedClient.GetUser()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
