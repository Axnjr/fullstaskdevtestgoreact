package models

import "time"

// Stock represents a stock with its current price
type Stock struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

// Order represents a buy/sell order
type Order struct {
	ID        int       `json:"id"`
	Symbol    string    `json:"symbol"`
	Side      string    `json:"side"` // "buy" or "sell"
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id,omitempty"` // User who placed the order
}

// PriceUpdate represents a price update message
type PriceUpdate struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Change    float64 `json:"change"`
	ChangePct float64 `json:"changePct"`
}

// User represents a user account
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Don't serialize password
}

// LoginRequest represents a login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents a login response with JWT token
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

