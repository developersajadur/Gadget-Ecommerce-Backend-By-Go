package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port      string
<<<<<<< HEAD
	Db_User_Name string
	Db_Password string
=======
	JwtSecret string
>>>>>>> 6e59bf6cae22451e31c634804160434ec6a8fda7
	// DBUrl     string
	// JWTSecret string
	// Env       string
}


var ENV *config


func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
		os.Exit(1)
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf(" Required environment variable %s is not set or empty", key)
		os.Exit(1)
	}
	return value
}


func Init() {
	loadEnv() 

	ENV = &config{
		Port: getEnv("PORT"),
<<<<<<< HEAD
		Db_User_Name: getEnv("DB_USER_NAME"),
		Db_Password: getEnv("DB_PASSWORD"),
=======
		JwtSecret: getEnv("JWT_SECRET"),
>>>>>>> 6e59bf6cae22451e31c634804160434ec6a8fda7
	}
}