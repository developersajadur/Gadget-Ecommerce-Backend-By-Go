package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Can't Decoded User From Body")
		return
	}

	user, err := h.userUC.Register(req.Name, req.Email, req.Password)
	if err != nil {
		fmt.Println(err)
		helpers.SendError(w, err, http.StatusConflict, "Already have a user by this email")
		return
	}
	helpers.SendResponse(w, user, http.StatusOK, "User created Successfully")
}
