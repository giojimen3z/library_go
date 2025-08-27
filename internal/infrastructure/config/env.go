package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found, reading environment variables from system", "error", err)
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func MustGetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	slog.Error("environment variable is required but not set", "key", key)
	os.Exit(1)
	return ""
}
