package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/domain"
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

	// Protected routes (role-based)
	r.Handle("/users/list", middleware.Middlewares(
		http.HandlerFunc(userHandler.List),
		middleware.Auth(uc, []string{domain.Role.Admin}),
	)).Methods("GET")

	// Public routes
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUserById).Methods("GET")
}
