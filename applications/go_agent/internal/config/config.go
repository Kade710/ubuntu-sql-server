package config

import (
    "fmt"
    "os"
)

//Config contains runtime settings for the Go agent.
type Config struct {
    DBHost string
    DBport string
    DBName string
    DBUser string
    DBPassword string
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
    cfg := Config{
        DBHost: getEnv("DB_HOST", "localhost"),
        DBPort: getEnv("DB_PORT", "5432"),
        DBName: os.Getenv("DB_NAME"),
        DBUser: os.Getenv("DB_USER"),
        DBPassword: os Getenv("DB_PASSWORD"),
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