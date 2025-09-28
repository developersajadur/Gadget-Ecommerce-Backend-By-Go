package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port      string
	// DBUrl     string
	// JWTSecret string
	// Env       string
}


var ENV *config


func loadEnv() {
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


func Init() {
	loadEnv() 

	ENV = &config{
		Port: getEnv("PORT"),
	}
}