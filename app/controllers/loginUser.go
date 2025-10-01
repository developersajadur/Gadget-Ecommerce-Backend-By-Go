package controllers

import (
	"ecommerce/app/config"
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"ecommerce/app/utils"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest, "Invalid request body")
		return
	}

	if creds.Email == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "Email is required")
		return
	}
	if creds.Password == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "Password is required")
		return
	}

	var user *database.Person
	for _, person := range database.People {
		if person.Email == creds.Email {
			user = &person
			break
		}
	}

	if user == nil {
		helpers.SendError(w, nil, http.StatusNotFound, "User not found")
		return
	}

	if user.Password != creds.Password {
		helpers.SendError(w, nil, http.StatusUnauthorized, "Password is incorrect")
		return
	}

	// Generate JWT
	payload := utils.JwtCustomClaims{
		UserId: user.ID,
		Name:   user.Name,
		Age:    user.Age,
		Email:  user.Email,
	}
	token, err := utils.GenerateJWT([]byte(config.ENV.JwtSecret), payload)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Optional: set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // set true in production
		Path:    "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24 * 3,
	})

	helpers.SendResponse(w, map[string]string{"token": token}, http.StatusOK, "Login successful")
}
