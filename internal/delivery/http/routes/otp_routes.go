package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterOtpRoutes(r *mux.Router, userUC usecase.UserUsecase, otpUc usecase.OtpUsecase) {

	otpHandler := handlers.NewOtpHandler(userUC, otpUc)

	r.HandleFunc("/send", otpHandler.Create).Methods("POST")
	r.HandleFunc("/verify", otpHandler.VerifyOtp).Methods("POST")
}
