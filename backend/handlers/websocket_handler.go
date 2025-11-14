package handlers

import (
	"log"
	"net/http"

	"trading-dashboard/models"
	"trading-dashboard/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

type WebSocketHandler struct {
	priceSimulator *services.PriceSimulator
}

func NewWebSocketHandler(priceSimulator *services.PriceSimulator) *WebSocketHandler {
	return &WebSocketHandler{
		priceSimulator: priceSimulator,
	}
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	// Create a channel for this client
	updateChan := make(chan models.PriceUpdate, 10)
	h.priceSimulator.Subscribe(updateChan)
	defer h.priceSimulator.Unsubscribe(updateChan)

	// Send initial prices
	stocks := h.priceSimulator.GetStocks()
	for _, stock := range stocks {
		update := models.PriceUpdate{
			Symbol:    stock.Symbol,
			Price:     stock.Price,
			Change:    0,
			ChangePct: 0,
		}
		if err := conn.WriteJSON(update); err != nil {
			log.Printf("Error sending initial price: %v", err)
			return
		}
	}

	// Listen for price updates and send to client
	go func() {
		for update := range updateChan {
			if err := conn.WriteJSON(update); err != nil {
				log.Printf("Error writing to WebSocket: %v", err)
				return
			}
		}
	}()
	
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}
	}
}
