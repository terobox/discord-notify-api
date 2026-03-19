package controller

import (
	"log"
	"main/config"
	"main/dto"
	"main/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	discordService *infra.DiscordService
}

func NewMessageHandler(cfg *config.Config) *MessageHandler {
	return &MessageHandler{
		discordService: infra.NewDiscordService(cfg),
	}
}

func (h *MessageHandler) Send(c *gin.Context) {
	var req dto.SendMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.SendMessageResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	if err := h.discordService.SendMessage(req); err != nil {
		log.Printf("❌ Failed to send message to channel %s: %v", req.ChannelID, err)
		c.JSON(http.StatusInternalServerError, dto.SendMessageResponse{
			Success: false,
			Message: "Failed to send message: " + err.Error(),
		})
		return
	}

	log.Printf("✅ Message sent to channel %s [source: %s]", req.ChannelID, req.Source)

	c.JSON(http.StatusOK, dto.SendMessageResponse{
		Success: true,
		Message: "Message sent successfully",
	})
}
