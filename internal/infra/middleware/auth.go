package middleware

import (
	"context"
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"ecommerce/pkg/utils/jwt"
	"errors"
	"net/http"
	"strconv"
)

type contextKey string

const UserContextKey = contextKey("user")

func Auth(userUC usecase.UserUsecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")
			if token == "" {
				helpers.SendError(w, nil, http.StatusUnauthorized, "Missing token")
				return
			}

			claims, err := jwt.GetUserFromJwt(token)
			if err != nil {
				helpers.SendError(w, err, http.StatusUnauthorized, "Invalid token")
				return
			}

			userIDVal, ok := claims["userId"]
			if !ok {
				helpers.SendError(w, nil, http.StatusUnauthorized, "Unauthorized")
				return
			}

			// Convert userID to string
			var userID string
			switch v := userIDVal.(type) {
			case float64:
				userID = strconv.FormatInt(int64(v), 10)
			case string:
				userID = v
			default:
				helpers.SendError(w, nil, http.StatusUnauthorized, "Invalid userId in token")
				return
			}

			user, err := userUC.GetUserById(userID)
			if err != nil || user == nil {
				helpers.SendError(w, errors.New("user not found"), http.StatusUnauthorized, "User not found")
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
