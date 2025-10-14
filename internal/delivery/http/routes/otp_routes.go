package routes

import (
	"ecommerce/internal/delivery/http/handlers"
	"ecommerce/internal/usecase"

	"github.com/gorilla/mux"
)

const (
	// OTP routes
	OtpRouteSend   = "/send"
	OtpRouteVerify = "/verify"
)

func RegisterOtpRoutes(r *mux.Router, userUC usecase.UserUsecase, otpUc usecase.OtpUsecase) {

	otpHandler := handlers.NewOtpHandler(userUC, otpUc)

	r.HandleFunc(OtpRouteSend, otpHandler.Create).Methods("POST")
	r.HandleFunc(OtpRouteVerify, otpHandler.VerifyOtp).Methods("POST")
}
