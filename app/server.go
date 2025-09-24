package utils

import (
	"fmt"
	"net/http"

	"ecommerce/app/routes"
)

func RunServer() {
	routes.SetupRoutes()

	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", routes.Mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
