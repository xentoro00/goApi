package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Port              string
	SupabaseURL       string
	SupabaseProjectID string // Add this new field
	SupabaseKey       string
	SecretAPIKey      string // Add this new field
}

// Load creates a Config struct from environment variables
func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:              os.Getenv("PORT"),
		SupabaseURL:       os.Getenv("SUPABASE_URL"),
		SupabaseProjectID: os.Getenv("SUPABASE_PROJECT_ID"), // Load the new variable
		SupabaseKey:       os.Getenv("SUPABASE_KEY"),
		SecretAPIKey:      os.Getenv("SECRET_API_KEY"),
	}

	return cfg, nil
}
