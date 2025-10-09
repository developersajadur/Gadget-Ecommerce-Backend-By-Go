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
	api_version_1 := Router.PathPrefix("/api/v1").Subrouter()

	Router.Use(middleware.Cors)

	RegisterUserRoutes(api_version_1, dbConn)

	return Router
}
