package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/usecase"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(mux *http.ServeMux, db *sqlx.DB) {
	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	userHandler := handlers.NewUserHandler(uc)

	mux.Handle("POST /users/create", http.HandlerFunc(userHandler.Register))
	mux.Handle("POST /users/login", http.HandlerFunc(userHandler.Login))
}
