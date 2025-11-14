package config

import (
	"os"
	"time"
)

// Config holds application configuration
type Config struct {
	JWTSecret     string
	JWTExpiration time.Duration
	Port          string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production" // Default for development
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		JWTSecret:     jwtSecret,
		JWTExpiration: 24 * time.Hour, // Token expires in 24 hours
		Port:          port,
	}
}

