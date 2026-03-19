package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"main/dto"
	"main/util"
	"net/http"
)

type DiscordService struct {
	botToken string
	client   *http.Client
}

func NewDiscordService(cfg *config.Config) *DiscordService {
	return &DiscordService{
		botToken: cfg.DiscordBotToken,
		client:   &http.Client{},
	}
}

func (s *DiscordService) SendMessage(req dto.SendMessageRequest) error {
	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", req.ChannelID)

	// Format the message with nice markdown
	formattedMsg := util.BuildDiscordMessage(req.Title, req.Content, req.Source)

	payload := map[string]string{
		"content": formattedMsg,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %v", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bot %s", s.botToken))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("discord API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
