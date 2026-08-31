package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config contains the configuration settings used by the Go agent.
type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	NTFYTopic  string
	APIToken   string
}

// Load reads configuration from environment variables and validates
// the required database settings.
func Load() (Config, error) {
	cfg := Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     strings.TrimSpace(os.Getenv("DB_NAME")),
		DBUser:     strings.TrimSpace(os.Getenv("DB_USER")),
		DBPassword: os.Getenv("DB_PASSWORD"),
		NTFYTopic:  strings.TrimSpace(os.Getenv("NTFY_TOPIC")),
		APIToken:	strings.TrimSpace(os.Getenv("API_TOKEN")),
	}

	if cfg.DBName == "" {
		return Config{}, fmt.Errorf("DB_NAME is required")
	}

	if cfg.DBUser == "" {
		return Config{}, fmt.Errorf("DB_USER is required")
	}

	if cfg.DBPassword == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is required")
	}

	port, err := strconv.Atoi(cfg.DBPort)
	if err != nil || port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("DB_PORT must be a valid TCP port")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}
