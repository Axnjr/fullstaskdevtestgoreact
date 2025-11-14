package services

import (
	"crypto/subtle"
	"errors"
	"sync"

	"trading-dashboard/models"
)

type AuthService struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

func NewAuthService() *AuthService {
	// Initialize with some default users (in production, use a database)
	users := map[string]*models.User{
		"admin": {
			ID:       "1",
			Username: "admin",
			Password: "admin123",
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

func (as *AuthService) Authenticate(username, password string) (*models.User, error) {
	as.mu.RLock()
	defer as.mu.RUnlock()

	user, exists := as.users[username]
	if !exists {
		return nil, errors.New("invalid credentials")
	}

	// Simple password comparison
	if subtle.ConstantTimeCompare([]byte(user.Password), []byte(password)) != 1 {
		return nil, errors.New("invalid credentials")
	}

	// Return user without password
	return &models.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

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

