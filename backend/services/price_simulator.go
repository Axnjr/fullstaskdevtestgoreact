package services

import (
	"math/rand"
	"sync"
	"time"

	"trading-dashboard/models"
)

// PriceSimulator manages stock prices and broadcasts updates
type PriceSimulator struct {
	stocks     map[string]*models.Stock
	stocksLock sync.RWMutex
	clients    map[chan models.PriceUpdate]bool
	clientsLock sync.RWMutex
}

// NewPriceSimulator creates a new price simulator
func NewPriceSimulator() *PriceSimulator {
	// Initialize with mock stocks
	stocks := map[string]*models.Stock{
		"AAPL": {Symbol: "AAPL", Price: 175.50},
		"TSLA": {Symbol: "TSLA", Price: 245.30},
		"AMZN": {Symbol: "AMZN", Price: 142.80},
		"INFY": {Symbol: "INFY", Price: 18.25},
		"TCS":  {Symbol: "TCS", Price: 3450.75},
	}

	return &PriceSimulator{
		stocks:  stocks,
		clients: make(map[chan models.PriceUpdate]bool),
	}
}

// GetStocks returns all stocks
func (ps *PriceSimulator) GetStocks() []models.Stock {
	ps.stocksLock.RLock()
	defer ps.stocksLock.RUnlock()

	stocks := make([]models.Stock, 0, len(ps.stocks))
	for _, stock := range ps.stocks {
		stocks = append(stocks, *stock)
	}
	return stocks
}

// GetStock returns a specific stock by symbol
func (ps *PriceSimulator) GetStock(symbol string) (*models.Stock, bool) {
	ps.stocksLock.RLock()
	defer ps.stocksLock.RUnlock()

	stock, exists := ps.stocks[symbol]
	if !exists {
		return nil, false
	}
	return stock, true
}

// Subscribe adds a channel to receive price updates
func (ps *PriceSimulator) Subscribe(ch chan models.PriceUpdate) {
	ps.clientsLock.Lock()
	defer ps.clientsLock.Unlock()
	ps.clients[ch] = true
}

// Unsubscribe removes a channel from price updates
func (ps *PriceSimulator) Unsubscribe(ch chan models.PriceUpdate) {
	ps.clientsLock.Lock()
	defer ps.clientsLock.Unlock()
	delete(ps.clients, ch)
	close(ch)
}

// broadcastUpdate sends price update to all subscribed channels
func (ps *PriceSimulator) broadcastUpdate(update models.PriceUpdate) {
	ps.clientsLock.RLock()
	defer ps.clientsLock.RUnlock()

	for ch := range ps.clients {
		select {
		case ch <- update:
		default:
			// Channel is full, skip this client
		}
	}
}

// Start begins the price simulation loop
func (ps *PriceSimulator) Start() {
	ticker := time.NewTicker(3 * time.Second) // Update every 3 seconds
	defer ticker.Stop()

	for range ticker.C {
		ps.stocksLock.Lock()
		for symbol, stock := range ps.stocks {
			oldPrice := stock.Price

			// Random change between -2% and +2%
			changePercent := (rand.Float64()*4 - 2) / 100 // -2% to +2%
			newPrice := oldPrice * (1 + changePercent)

			// Ensure price doesn't go below 0.01
			if newPrice < 0.01 {
				newPrice = 0.01
			}

			stock.Price = newPrice

			// Broadcast update
			update := models.PriceUpdate{
				Symbol:    symbol,
				Price:     newPrice,
				Change:    newPrice - oldPrice,
				ChangePct: changePercent * 100,
			}

			ps.stocksLock.Unlock()
			ps.broadcastUpdate(update)
			ps.stocksLock.Lock()
		}
		ps.stocksLock.Unlock()
	}
}

