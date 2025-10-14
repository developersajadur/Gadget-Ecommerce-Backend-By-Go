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
	RouteSlug   = "/slug/{slug}"
	RouteID     = "/id/{id}"
	RouteList   = "/list"
	RouteCreate = "/create"
	RouteUpdate = "/update/{id}"
	RouteDelete = "/delete/{id}"
)

func RegisterCategoryRoutes(r *mux.Router, categoryUC usecase.CategoryUsecase, userUC usecase.UserUsecase) {
	categoryHandler := handlers.NewCategoryHandler(categoryUC)

	// Public routes
	r.HandleFunc(RouteSlug, categoryHandler.GetBySlug).Methods("GET")
	r.HandleFunc(RouteID, categoryHandler.GetById).Methods("GET")
	r.HandleFunc(RouteList, categoryHandler.List).Methods("GET")

	// Admin routes
	r.Handle(RouteCreate, middleware.Middlewares(
		http.HandlerFunc(categoryHandler.Create),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("POST")

	r.Handle(RouteUpdate, middleware.Middlewares(
		http.HandlerFunc(categoryHandler.Update),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("PATCH")

	r.Handle(RouteDelete, middleware.Middlewares(
		http.HandlerFunc(categoryHandler.SoftDelete),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("DELETE")
}
