package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// RegisterOtpRoutes sets up OTP-related endpoints
func RegisterOtpRoutes(r *mux.Router, db *sqlx.DB) {
	otpRepo := repository.NewOtpRepository(db)

	// Usecases
	otpUC := usecase.NewOtpUsecase(otpRepo)

	// Handler
	otpHandler := handlers.NewOtpHandler(otpUC)

	// OTP routes
	// Send OTP for login or verification
	r.HandleFunc("/otp/send", otpHandler.Create).Methods("POST")

	// Verify OTP
	// r.HandleFunc("/otp/verify", userHandler.VerifyOTP).Methods("POST")
}
