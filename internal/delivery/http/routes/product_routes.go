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
	// Product routes
	ProductRouteSlug   = "/slug/{slug}"
	ProductRouteID     = "/id/{id}"
	ProductRouteList   = "/list"
	ProductRouteCreate = "/create"
	ProductRouteUpdate = "/update/{id}"
	ProductRouteDelete = "/delete/{id}"
)

func RegisterProductRoutes(r *mux.Router, productUC usecase.ProductUsecase, userUC usecase.UserUsecase) {
	productHandler := handlers.NewProductHandler(productUC)

	// Public Routes
	r.HandleFunc(ProductRouteSlug, productHandler.GetBySlug).Methods("GET")
	r.HandleFunc(ProductRouteID, productHandler.GetById).Methods("GET")
	r.HandleFunc(ProductRouteList, productHandler.List).Methods("GET")

	// Private Routes (Admin only)
	r.Handle(ProductRouteCreate, middleware.Middlewares(
		http.HandlerFunc(productHandler.Create),
		middleware.Auth(userUC, []string{models.RoleAdmin}),
	)).Methods("POST")

	// r.Handle(ProductRouteUpdate, middleware.Middlewares(
	// 	http.HandlerFunc(productHandler.Update),
	// 	middleware.Auth(userUC, []string{models.RoleAdmin}),
	// )).Methods("PATCH")

	// r.Handle(ProductRouteDelete, middleware.Middlewares(
	// 	http.HandlerFunc(productHandler.SoftDelete),
	// 	middleware.Auth(userUC, []string{models.RoleAdmin}),
	// )).Methods("DELETE")
}
