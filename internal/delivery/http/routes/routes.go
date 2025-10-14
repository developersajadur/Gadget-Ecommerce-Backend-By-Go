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
dbConn, err := db.ConnectDB()
if err != nil {
	panic(err)
}


	Router = mux.NewRouter()
	Router.Use(middleware.Cors)

	// Initialize Repositories
	userRepo := repository.NewUserRepository(dbConn)
	otpRepo := repository.NewOtpRepository(dbConn)
	categoryRepo := repository.NewCategoryRepository(dbConn)
	productRepo := repository.NewProductRepository(dbConn)

	// Initialize Usecases
	otpUC := usecase.NewOtpUsecase(otpRepo)
	userUC := usecase.NewUserUsecase(userRepo, otpUC)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo)
	productUC := usecase.NewProductUsecase(productRepo, categoryUC)

	// API Versioning
	apiV1 := Router.PathPrefix("/api/v1").Subrouter()

	// Subrouters for resources
	userRouter := apiV1.PathPrefix("/users").Subrouter()
	otpRouter := apiV1.PathPrefix("/otps").Subrouter()
	categoryRouter := apiV1.PathPrefix("/categories").Subrouter()
	productRouter := apiV1.PathPrefix("/products").Subrouter()

	// Register routes
	RegisterUserRoutes(userRouter, userUC)
	RegisterOtpRoutes(otpRouter, userUC, otpUC)
	RegisterCategoryRoutes(categoryRouter, categoryUC, userUC)
	RegisterProductRoutes(productRouter, productUC, userUC)

	return Router
}
