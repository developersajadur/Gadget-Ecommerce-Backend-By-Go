package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"net/http"
)

type OtpHandler struct {
	otpUC  usecase.OtpUsecase
	userUc usecase.UserUsecase
}

// Constructor now takes both userUC and otpUC
func NewOtpHandler(userUC usecase.UserUsecase, otpUC usecase.OtpUsecase) *OtpHandler {
	return &OtpHandler{
		userUc: userUC,
		otpUC:  otpUC,
	}
}

func (h *OtpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}

	helpers.BodyDecoder(w, r, &req)

	// 1. Find user
	user, err := h.userUc.FindByEmail(req.Email)
	if err != nil || user == nil {
		helpers.SendError(w, err, http.StatusNotFound, "User not found")
		return
	}

	// 2. Create OTP and send email
	otpEntry, err := h.otpUC.CreateAndSendEmail(user.ID, user.Name, user.Email)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to create OTP")
		return
	}

	helpers.SendResponse(w, map[string]string{
		"message": "OTP created and email sent successfully",
		"otp_id":  otpEntry.ID,
	}, http.StatusOK, "")
}
