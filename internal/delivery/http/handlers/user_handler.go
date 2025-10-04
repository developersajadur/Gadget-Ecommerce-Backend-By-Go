package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

// Register
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Can't decode user from body")
		return
	}

	user, err := h.userUC.Register(req.Name, req.Email, req.Password)
	if err != nil {
		helpers.SendError(w, err, http.StatusConflict, "User already exists with this email")
		return
	}

	helpers.SendResponse(w, user, http.StatusOK, "User created successfully")
}

// Login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Can't decode user from body")
		return
	}

	token, err := h.userUC.Login(req.Email, req.Password)
	if err != nil {
		helpers.SendError(w, err, http.StatusUnauthorized, "Invalid credentials or user not found")
		return
	}

	helpers.SendResponse(w, map[string]string{"token": token}, http.StatusOK, "Login successful")
}

