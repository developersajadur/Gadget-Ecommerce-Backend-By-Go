package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

// Create
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	helpers.BodyDecoder(w, r, &req)

	user, err := h.userUC.Create(req.Name, req.Email, req.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			helpers.SendError(w, err, http.StatusConflict, "User already exists with this email")
			return
		}
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to create user")
		return
	}

	helpers.SendResponse(w, user, http.StatusOK, "User created successfully")
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	users, err := h.userUC.List(page, limit)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	helpers.SendResponse(w, users, http.StatusOK, "Login successful")

}

// Login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	helpers.BodyDecoder(w, r, &req)

	token, err := h.userUC.Login(req.Email, req.Password)
	if err != nil {
		helpers.SendError(w, err, http.StatusUnauthorized, "Invalid credentials or user not found")
		return
	}

	helpers.SendResponse(w, map[string]string{"token": token}, http.StatusOK, "Login successful")
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.userUC.GetUserById(id)
	if err != nil {
		helpers.SendError(w, err, http.StatusNotFound, "User not found")
		return
	}

	helpers.SendResponse(w, user, http.StatusOK, "User fetched successfully")
}
