package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"ecommerce/pkg/utils/jwt"
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
	search := r.URL.Query().Get("search")
	users, err := h.userUC.List(page, limit, search)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	helpers.SendResponse(w, users, http.StatusOK, "Users fetched successfully")

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
		var status int
		switch err.Error() {
		case "user not found":
			status = http.StatusNotFound
		case "user is blocked", "user is not verified":
			status = http.StatusForbidden
		case "invalid credentials":
			status = http.StatusUnauthorized
		default:
			status = http.StatusInternalServerError
		}
		helpers.SendError(w, err, status, err.Error())
		return
	}

	helpers.SendResponse(w, map[string]string{"token": token}, http.StatusOK, "")
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

func (h *UserHandler) GetMyUserDetails(w http.ResponseWriter, r *http.Request) {

	jwtUser, err := jwt.GetUserFromJwt(r)
	if err != nil {
		helpers.SendError(w, err, http.StatusUnauthorized, "Invalid token")
		return
	}

	user, err := h.userUC.GetUserById(jwtUser.UserID)
	if err != nil {
		helpers.SendError(w, err, http.StatusNotFound, "User not found")
		return
	}

	helpers.SendResponse(w, user, http.StatusOK, "User fetched successfully")
}
func (h *UserHandler) BlockUserByAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.userUC.BlockUserByAdmin(id)
	if err != nil {
		switch err.Error() {
		case "user not found":
			helpers.SendError(w, err, http.StatusNotFound, err.Error())
		case "user is already blocked":
			helpers.SendError(w, err, http.StatusBadRequest, err.Error())
		default:
			helpers.SendError(w, err, http.StatusInternalServerError, "Failed to block user")
		}
		return
	}

	helpers.SendResponse(w, nil, http.StatusOK, "User blocked successfully")
}


func (h *UserHandler) UnblockUserByAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.userUC.UnblockUserByAdmin(id)
	if err != nil {
		switch err.Error() {
		case "user not found":
			helpers.SendError(w, err, http.StatusNotFound, err.Error())
		case "user is already unblocked":
			helpers.SendError(w, err, http.StatusBadRequest, err.Error())
		default:
			helpers.SendError(w, err, http.StatusInternalServerError, "Failed to unBlock user")
		}
		return
	}

	helpers.SendResponse(w, nil, http.StatusOK, "User unblocked successfully")
}
