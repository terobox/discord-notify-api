package main

import (
	"log"
	"main/config"
	"main/router"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}

	// Set Gin mode
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Set routes
	router.SetRouter(r, cfg)

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Discord Notify API starting on :%s", port)
	log.Printf("📡 Environment: %s", cfg.Env)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
