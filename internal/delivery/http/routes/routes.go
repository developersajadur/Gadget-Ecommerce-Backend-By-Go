package routes

import (
	"ecommerce/internal/infra/db"
	"ecommerce/internal/infra/middleware"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func SetupRoutes() *mux.Router {
	dbConn, _ := db.NewConnection()
	Router = mux.NewRouter()

	Router.Use(middleware.Cors)

	RegisterUserRoutes(Router, dbConn)

	return Router
}
