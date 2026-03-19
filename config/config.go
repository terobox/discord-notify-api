package config

import (
	"fmt"
	"os"
)

type Config struct {
	// Server
	Port string
	Env  string

	// Auth
	APIKey string

	// Discord
	DiscordBotToken string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:            getEnv("PORT", "8080"),
		Env:             getEnv("ENV", "development"),
		APIKey:          os.Getenv("API_KEY"),
		DiscordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("API_KEY is required (should start with 'sk-')")
	}

	if cfg.DiscordBotToken == "" {
		return nil, fmt.Errorf("DISCORD_BOT_TOKEN is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
