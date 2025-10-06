package main

import (
	"ecommerce/internal/config"
	"ecommerce/internal/delivery/http/routes"
	"fmt"
	"net/http"
)

func RunServer() {

	routes.SetupRoutes()

	fmt.Println("Starting Server At:", config.ENV.Port)
	if err := http.ListenAndServe(":"+config.ENV.Port, routes.Router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
