package utils

import (
	"fmt"
	"net/http"

	"ecommerce/app/config"
	"ecommerce/app/routes"
)

func RunServer() {
	config.Init()
	routes.SetupRoutes()
	fmt.Println("Starting server at:", config.ENV.Port)
	err := http.ListenAndServe(":"+config.ENV.Port, routes.Mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		
	}
}
