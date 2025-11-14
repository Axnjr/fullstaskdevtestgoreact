package services

import (
	"context"
	"sync"
	"time"

	"trading-dashboard/models"
)

var ctx = context.Background()

type OrderService struct {
	orders     []models.Order
	ordersLock sync.RWMutex
	orderID    int
}

func NewOrderService() *OrderService {
	return &OrderService{
		orders:  make([]models.Order, 0),
		orderID: 1,
	}
}

// storing these orders in memory, persistence does not makes sense for this demo :)
func (os *OrderService) CreateOrder(order models.Order) models.Order {
	os.ordersLock.Lock()
	defer os.ordersLock.Unlock()

	order.ID = os.orderID
	os.orderID++
	order.Timestamp = time.Now()
	os.orders = append(os.orders, order)

	return order
}

func (os *OrderService) GetOrders() []models.Order {
	os.ordersLock.RLock()
	defer os.ordersLock.RUnlock()

	// Return a copy to prevent external modification
	orders := make([]models.Order, len(os.orders))
	copy(orders, os.orders)
	return orders
}

func (os *OrderService) GetOrdersByUser(userID string) []models.Order {
	os.ordersLock.RLock()
	defer os.ordersLock.RUnlock()

	var userOrders []models.Order
	for _, order := range os.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}

