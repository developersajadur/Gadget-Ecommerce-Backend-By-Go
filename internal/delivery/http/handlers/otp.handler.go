package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"net/http"
)




type OtpHandler struct {
	otpUC usecase.OtpUsecase
}

func NewOtpHandler(uc usecase.OtpUsecase) *OtpHandler {
	return &OtpHandler{uc}
}

func (h *OtpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID string `json:"user_id"`
	}

	helpers.BodyDecoder(w, r, &req)

	otpEntry, err := h.otpUC.Create(req.UserID)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to create OTP")
		return
	}

	helpers.SendResponse(w, otpEntry, http.StatusOK, "OTP created successfully")
}