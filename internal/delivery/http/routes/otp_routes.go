package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RegisterOtpRoutes(r *mux.Router, db *sqlx.DB, userUC usecase.UserUsecase) {
	otpRepo := repository.NewOtpRepository(db)
	otpUC := usecase.NewOtpUsecase(otpRepo)
	otpHandler := handlers.NewOtpHandler(userUC, otpUC)

	r.HandleFunc("/send", otpHandler.Create).Methods("POST")
}
