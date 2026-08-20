package config

import (
	"fmt"
	"os"
)

// Config contains the PostgreSQL connection settings.
type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	NTFYTopic  string
}

// Load reads PostgreSQL configuration from environment variables.
func Load() (Config, error) {
	cfg := Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		NTFYTopic:  os.Getenv("NTFY_TOPIC"),
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

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
