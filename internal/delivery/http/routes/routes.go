package routes

import (
	"ecommerce/internal/infra/db"
	"ecommerce/internal/infra/middleware"
	"net/http"
)

var Mux *http.ServeMux

func SetupRoutes() {
	db, _ := db.NewConnection()
	Mux = http.NewServeMux()

	RegisterUserRoutes(Mux, db)

	middleware.Cors(Mux)
	return

}
