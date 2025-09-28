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

	for _, person := range database.People {
		if person.Email != creds.Email {
			helpers.SendError(w, nil, http.StatusNotFound, "User Not Found By This Email")
			return
		} else if person.Password != creds.Password {
			helpers.SendError(w, nil, http.StatusUnauthorized, "Password is Incorrect")
			return
		}
		payload := utils.JwtCustomClaims{
			UserId: person.ID,
			Name:   person.Name,
			Age:    person.Age,
			Email:  person.Email,
		}
		jwtSecret := []byte(config.ENV.JwtSecret)
		token, err := utils.GenerateJWT(jwtSecret, payload)
		if err != nil {
			helpers.SendError(w, err, http.StatusInternalServerError, "Failed to generate token")
			return
		}
		helpers.SendResponse(w, map[string]string{"token": token}, http.StatusOK, "Login successful")
		return

	}

}
