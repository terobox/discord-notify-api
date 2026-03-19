package router

import (
	"main/config"
	"main/controller"
	"main/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine, cfg *config.Config) {
	// Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Discord Notify API!",
		})
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthRequired(cfg))
	{
		msgHandler := controller.NewMessageHandler(cfg)
		v1.POST("/send", msgHandler.Send)
	}
}
