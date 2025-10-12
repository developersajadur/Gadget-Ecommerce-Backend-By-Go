package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/middleware"
	"ecommerce/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(r *mux.Router, categoryUC usecase.CategoryUsecase, userUC usecase.UserUsecase) {

	categoryHandler := handlers.NewCategoryHandler(categoryUC)

	// Private Routes
	r.Handle("/create", middleware.Middlewares(
		http.HandlerFunc(categoryHandler.Create),
		middleware.Auth(userUC, []string{domain.RoleAdmin}),
	)).Methods("POST")

}
