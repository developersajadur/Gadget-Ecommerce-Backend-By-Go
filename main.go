package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/routes"
	"ecommerce/internal/infra/db"
	"ecommerce/internal/infra/middleware"
	"ecommerce/pkg/utils/email"
)

func init() {
	config.Init()

	db.ConnectDB()

	db.RunMigrations()
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

	router := routes.SetupRoutes()

	handler := middleware.NewRateLimiterMiddleware(middleware.RateLimiterConfig{
		Limit:  4,
		Period: 1 * time.Second,
	})(router)

	fmt.Printf("Server running at port %s\n", config.ENV.Port)
	if err := http.ListenAndServe(":"+config.ENV.Port, handler); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	runServer()
}
