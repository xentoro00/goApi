package main

import (
	config "go-messaging-api/internal"
	"go-messaging-api/internal/api"
	"log"

	supa "github.com/nedpals/supabase-go"
	supaauth "github.com/supabase-community/auth-go"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize BOTH Supabase clients with the correct variables
	// FIX: Use the Project ID for the auth client
	authClient := supaauth.New(cfg.SupabaseProjectID, cfg.SupabaseKey)

	// FIX: Use the full URL for the database client
	dbClient := supa.CreateClient(cfg.SupabaseURL, cfg.SupabaseKey)

	// Pass both clients to the router
	router := api.NewRouter(&authClient, dbClient)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
