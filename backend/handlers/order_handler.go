package handlers

import (
	"net/http"

	"trading-dashboard/models"
	"trading-dashboard/services"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles order-related requests
type OrderHandler struct {
	orderService   *services.OrderService
	priceSimulator *services.PriceSimulator
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(orderService *services.OrderService, priceSimulator *services.PriceSimulator) *OrderHandler {
	return &OrderHandler{
		orderService:   orderService,
		priceSimulator: priceSimulator,
	}
}

// CreateOrder handles POST /orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if exists {
		order.UserID = userID.(string)
	}

	// Validate order
	if order.Side != "buy" && order.Side != "sell" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "side must be 'buy' or 'sell'"})
		return
	}

	if order.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "quantity must be greater than 0"})
		return
	}

	if order.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price must be greater than 0"})
		return
	}

	// Validate symbol exists
	_, exists = h.priceSimulator.GetStock(order.Symbol)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid symbol"})
		return
	}

	// Create order
	createdOrder := h.orderService.CreateOrder(order)
	c.JSON(http.StatusCreated, createdOrder)
}

// GetOrders handles GET /orders
func (h *OrderHandler) GetOrders(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")

	var orders []models.Order
	if exists {
		// Return only user's orders
		orders = h.orderService.GetOrdersByUser(userID.(string))
	} else {
		// Fallback: return all orders (shouldn't happen if middleware is applied)
		orders = h.orderService.GetOrders()
	}

	c.JSON(http.StatusOK, orders)
}
