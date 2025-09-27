package utils

import (
	"fmt"
	"net/http"

	"ecommerce/app/config"
	"ecommerce/app/routes"
)

func RunServer() {
	cfg := config.NewConfig()

	routes.SetupRoutes()
	fmt.Println("Starting server at:", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, routes.Mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		
	}
}
