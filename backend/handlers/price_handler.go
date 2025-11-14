package handlers

import (
	"trading-dashboard/services"

	"github.com/gin-gonic/gin"
)

// PriceHandler handles price-related requests
type PriceHandler struct {
	priceSimulator *services.PriceSimulator
}

// NewPriceHandler creates a new price handler
func NewPriceHandler(priceSimulator *services.PriceSimulator) *PriceHandler {
	return &PriceHandler{
		priceSimulator: priceSimulator,
	}
}

// GetPrices handles GET /prices
func (h *PriceHandler) GetPrices(c *gin.Context) {
	stocks := h.priceSimulator.GetStocks()
	c.JSON(200, stocks)
}
