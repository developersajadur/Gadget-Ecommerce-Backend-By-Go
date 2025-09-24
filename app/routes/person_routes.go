package routes

import (
	"net/http"
	"ecommerce/app/controllers"
	"ecommerce/app/middlewares"
)

var Mux = http.NewServeMux()

func SetupRoutes() {
	Mux.Handle("GET /", http.HandlerFunc(controllers.RootPath))
	Mux.Handle("GET /persons", middleware.Middleware(http.HandlerFunc(controllers.GetPersons)))
	Mux.Handle("POST /persons/create", middleware.Middleware(http.HandlerFunc(controllers.CreatePerson)))
}
