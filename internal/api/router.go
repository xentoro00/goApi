package api

import (
	"go-messaging-api/internal/api/middleware"
	"go-messaging-api/internal/auth"
	"go-messaging-api/internal/members"
	"go-messaging-api/internal/rooms"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
	supaauth "github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

// FIX: NewRouter now accepts both the auth client and the database client
func NewRouter(authClient *supaauth.Client, dbClient *supa.Client) *gin.Engine {
	router := gin.Default()
	// --- CORS Middleware ---
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// --- NEW: Origin Check Middleware ---
	router.Use(middleware.CheckOrigin())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Setup auth routes
	authSvc := auth.NewService(*authClient)
	authHandler := auth.NewHandler(authSvc)
	authGroup := router.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	// Setup protected routes
	v1 := router.Group("/v1")
	v1.Use(middleware.Protected(*authClient))
	{
		v1.GET("/profile", profileHandler)

		// Setup rooms routes
		roomRepo := rooms.NewRepository(dbClient)
		roomSvc := rooms.NewService(roomRepo)
		roomHandler := rooms.NewHandler(roomSvc)
		roomHandler.RegisterRoutes(v1)

		// Setup members routes
		memberRepo := members.NewRepository(dbClient)
		memberSvc := members.NewService(memberRepo)
		memberHandler := members.NewHandler(memberSvc)
		memberHandler.RegisterRoutes(v1)
	}

	return router
}
func profileHandler(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}

	// FIX: Assert the type as a VALUE here as well
	_, ok := user.(types.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user data in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "welcome to your profile", "user": user})
}
