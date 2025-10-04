package main

import (
	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/routes"
	migrateDb "ecommerce/internal/infra/db"
	"fmt"
	"net/http"
	"os"
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

	routes.SetupRoutes()

	fmt.Println("Starting Server At:", config.ENV.Port)
	if err := http.ListenAndServe(":"+config.ENV.Port, routes.Mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	runServer()
}
