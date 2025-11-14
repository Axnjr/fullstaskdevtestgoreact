package config

import (
	"os"
	"time"
)

type Config struct {
	JWTSecret     string
	JWTExpiration time.Duration
	Port          string
}

func LoadConfig() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-change-in-production"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		JWTSecret:     jwtSecret,
		JWTExpiration: 24 * time.Hour, // expires in 24 hours
		Port:          port,
	}
}

