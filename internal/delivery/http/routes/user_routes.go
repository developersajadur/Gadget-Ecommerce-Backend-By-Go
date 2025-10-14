package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/infra/middleware"
	"ecommerce/internal/models"
	"ecommerce/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// User routes
	UserRouteList           = "/list"
	UserRouteGetByID        = "/list/{id}"
	UserRouteGetMyDetails   = "/get-my-user-details"
	UserRouteBlockUser      = "/block-user/{id}"
	UserRouteUnblockUser    = "/unblock-user/{id}"
	UserRouteCreate         = "/create"
	UserRouteLogin          = "/login"
)

func RegisterUserRoutes(r *mux.Router, userUC usecase.UserUsecase) {
	userHandler := handlers.NewUserHandler(userUC)

	// Protected routes (role-based)
	r.Handle(UserRouteList, middleware.Middlewares(
		http.HandlerFunc(userHandler.List),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("GET")

	r.Handle(UserRouteGetByID, middleware.Middlewares(
		http.HandlerFunc(userHandler.GetUserById),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("GET")

	r.Handle(UserRouteGetMyDetails, middleware.Middlewares(
		http.HandlerFunc(userHandler.GetMyUserDetails),
		middleware.Auth(userUC, []string{models.RoleAdmin, models.RoleUser}),
	)).Methods("GET")

	r.Handle(UserRouteBlockUser, middleware.Middlewares(
		http.HandlerFunc(userHandler.BlockUserByAdmin),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("POST")

	r.Handle(UserRouteUnblockUser, middleware.Middlewares(
		http.HandlerFunc(userHandler.UnblockUserByAdmin),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("POST")

	// Public routes
	r.HandleFunc(UserRouteCreate, userHandler.Create).Methods("POST")
	r.HandleFunc(UserRouteLogin, userHandler.Login).Methods("POST")
}
