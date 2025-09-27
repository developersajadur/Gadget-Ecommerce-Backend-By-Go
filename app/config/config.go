package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	// DBUrl     string
	// JWTSecret string
	// Env       string
}


func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found. Using system environment variables.")
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("❌ Required environment variable %s is not set or empty", key)
	}
	return value
}


func NewConfig() *Config {
	LoadEnv()

	cfg := &Config{
		Port:      getEnv("PORT"),
		// DBUrl:     getEnv("DATABASE_URL"),
		// JWTSecret: getEnv("JWT_SECRET"),
		// Env:       getEnv("APP_ENV"),
	}

	return cfg
}
