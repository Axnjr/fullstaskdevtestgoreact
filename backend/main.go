package main

import (
	"log"
	"math/rand"
	"time"

	"trading-dashboard/config"
	"trading-dashboard/handlers"
	"trading-dashboard/middleware"
	"trading-dashboard/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config.LoadConfig()

	authService := services.NewAuthService()
	orderService := services.NewOrderService()
	priceSimulator := services.NewPriceSimulator()

	authHandler := handlers.NewAuthHandler(authService, cfg)
	priceHandler := handlers.NewPriceHandler(priceSimulator)
	orderHandler := handlers.NewOrderHandler(orderService, priceSimulator)
	wsHandler := handlers.NewWebSocketHandler(priceSimulator)

	go priceSimulator.Start()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "https://fullstaskdevtestgoreact.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/login", authHandler.Login)
	r.GET("/prices", priceHandler.GetPrices)
	r.GET("/ws", wsHandler.HandleWebSocket)

	// Protected routes (require JWT authentication)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.POST("/orders", orderHandler.CreateOrder)
		protected.GET("/orders", orderHandler.GetOrders)
	}

	addr := "0.0.0.0:" + cfg.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
