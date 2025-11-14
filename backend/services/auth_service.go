package services

import (
	"crypto/subtle"
	"errors"
	"sync"

	"trading-dashboard/models"
)

// AuthService manages user authentication
type AuthService struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

// NewAuthService creates a new auth service with default users
func NewAuthService() *AuthService {
	// Initialize with some default users (in production, use a database)
	users := map[string]*models.User{
		"admin": {
			ID:       "1",
			Username: "admin",
			Password: "admin123", // In production, use bcrypt hashed passwords
		},
		"trader": {
			ID:       "2",
			Username: "trader",
			Password: "trader123",
		},
	}

	return &AuthService{
		users: users,
	}
}

// Authenticate validates username and password
func (as *AuthService) Authenticate(username, password string) (*models.User, error) {
	as.mu.RLock()
	defer as.mu.RUnlock()

	user, exists := as.users[username]
	if !exists {
		return nil, errors.New("invalid credentials")
	}

	// Simple password comparison (in production, use bcrypt)
	if subtle.ConstantTimeCompare([]byte(user.Password), []byte(password)) != 1 {
		return nil, errors.New("invalid credentials")
	}

	// Return user without password
	return &models.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

// GetUserByID retrieves a user by ID
func (as *AuthService) GetUserByID(userID string) (*models.User, error) {
	as.mu.RLock()
	defer as.mu.RUnlock()

	for _, user := range as.users {
		if user.ID == userID {
			return &models.User{
				ID:       user.ID,
				Username: user.Username,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}

