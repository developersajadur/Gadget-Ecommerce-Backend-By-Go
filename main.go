package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/routes"
	migrateDb "ecommerce/internal/infra/db"
	"ecommerce/internal/infra/middleware"
	"ecommerce/pkg/utils/email"
)

func init() {
	config.Init()

	db, err := migrateDb.NewConnection()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	fmt.Println("Database connected successfully")
	migrateDb.RunMigrations(db)
}

func runServer() {
	// Initialize email
	port, err := strconv.Atoi(config.ENV.Email_Port)
	if err != nil {
		fmt.Println("Invalid email port:", err)
		os.Exit(1)
	}

	email.Init(email.SMTPConfig{
		Host:     config.ENV.Email_Host,
		Port:     port,
		Username: config.ENV.Email,
		Password: config.ENV.Email_App_Password,
		From:     config.ENV.Email,
	})

	// Setup routes
	router := routes.SetupRoutes()

	// Rate limiter: 4 requests per second globally
	rateLimiter := middleware.NewRateLimiterMiddleware(middleware.RateLimiterConfig{
		Limit:  4,
		Period: 1 * time.Second,
	})

	handler := rateLimiter(router)

	fmt.Println("Starting Server At:", config.ENV.Port)
	if err := http.ListenAndServe(":"+config.ENV.Port, handler); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	runServer()
}
