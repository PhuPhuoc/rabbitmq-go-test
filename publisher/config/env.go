package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT        string
	RABBIT_HOST     string
	RABBIT_PORT     string
	RABBIT_USERNAME string
	RABBIT_PASSWORD string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		APP_PORT: getEnv("PORT", "8000"),
		// amqp://guest:guest@localhost:5672
		RABBIT_HOST:     getEnv("MONGODB_URI", "localhost"),
		RABBIT_PORT:     getEnv("MONGODB_DATABASE", "5672"),
		RABBIT_USERNAME: getEnv("RABBIT_USERNAME", "guest"),
		RABBIT_PASSWORD: getEnv("RABBIT_PASSWORD", "guest"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
