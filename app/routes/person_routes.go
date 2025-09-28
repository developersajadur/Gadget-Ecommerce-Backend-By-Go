package routes

import (
	"ecommerce/app/controllers"
	middleware "ecommerce/app/middlewares"
	"net/http"
)

var Mux = http.NewServeMux()

func SetupRoutes() {
	Mux.Handle("GET /", http.HandlerFunc(controllers.RootPath))
	Mux.Handle("GET /persons", middleware.Middlewares(
		http.HandlerFunc(controllers.GetPersons),
		middleware.Auth,
		middleware.AccessControl,
	))

	Mux.Handle("GET /persons/{personId}", middleware.Middlewares(
		http.HandlerFunc(controllers.GetPersonById),
		middleware.Auth,
		middleware.AccessControl,
	))
	Mux.Handle("POST /persons/create", middleware.Middlewares(
		http.HandlerFunc(controllers.CreatePerson),
		middleware.Auth,
		middleware.AccessControl,
	))
	Mux.Handle("POST /auth/login", middleware.Middlewares(
		http.HandlerFunc(controllers.LoginUser),
		middleware.AccessControl,
	))
}
