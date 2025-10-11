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
	userRepo := repository.NewUserRepository(db)
	otpRepo := repository.NewOtpRepository(db)

	otpUC := usecase.NewOtpUsecase(otpRepo)
	userUC := usecase.NewUserUsecase(userRepo, otpUC)

	userHandler := handlers.NewUserHandler(userUC)

	// Protected routes (role-based)
	r.Handle("/list", middleware.Middlewares(
		http.HandlerFunc(userHandler.List),
		middleware.Auth(userUC, []string{domain.RoleAdmin}),
	)).Methods("GET")

	r.Handle("/list/{id}", middleware.Middlewares(
		http.HandlerFunc(userHandler.GetUserById),
		middleware.Auth(userUC, []string{domain.RoleAdmin}),
	)).Methods("GET")

	r.Handle("/get-my-user-details", middleware.Middlewares(
		http.HandlerFunc(userHandler.GetMyUserDetails),
		middleware.Auth(userUC, []string{domain.RoleAdmin, domain.RoleUser}),
	)).Methods("GET")

	r.Handle("/block-user/{id}", middleware.Middlewares(
		http.HandlerFunc(userHandler.BlockUserByAdmin),
		middleware.Auth(userUC, []string{domain.RoleAdmin}),
	)).Methods("POST")

	r.Handle("/unblock-user/{id}", middleware.Middlewares(
		http.HandlerFunc(userHandler.UnblockUserByAdmin),
		middleware.Auth(userUC, []string{domain.RoleAdmin}),
	)).Methods("POST")

	// Public routes
	r.HandleFunc("/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
}

