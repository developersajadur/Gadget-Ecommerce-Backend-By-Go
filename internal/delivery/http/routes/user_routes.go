package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/infra/middleware"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(r *mux.Router, db *sqlx.DB) {
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	userHandler := handlers.NewUserHandler(uc)

	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")

	r.Handle("/users/list", middleware.Middlewares(
		http.HandlerFunc(userHandler.List),
		middleware.Auth(uc),
	)).Methods("GET")

	r.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUserById).Methods("GET")
}
