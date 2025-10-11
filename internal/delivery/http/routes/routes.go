package routes

import (
	"ecommerce/internal/infra/db"
	"ecommerce/internal/infra/middleware"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/usecase"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func SetupRoutes() *mux.Router {
	dbConn, _ := db.NewConnection()
	Router = mux.NewRouter()

	Router.Use(middleware.Cors)

	// Initialize Usecases
	userRepo := repository.NewUserRepository(dbConn)
	otpRepo := repository.NewOtpRepository(dbConn)

	otpUC := usecase.NewOtpUsecase(otpRepo)
	userUC := usecase.NewUserUsecase(userRepo, otpUC)

	// Versioned API
	apiV1 := Router.PathPrefix("/api/v1").Subrouter()

	// Create subrouters for specific resources
	userRouter := apiV1.PathPrefix("/users").Subrouter()
	otpRouter := apiV1.PathPrefix("/otps").Subrouter()

	// Register routes
	RegisterUserRoutes(userRouter, userUC)
	RegisterOtpRoutes(otpRouter, userUC, otpUC)

	return Router
}
