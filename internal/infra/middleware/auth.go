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


func Auth(userUC usecase.UserUsecase, roles []string) func(http.Handler) http.Handler {

	allowedRoles := make(map[string]bool)
	for _, r := range roles {
		allowedRoles[r] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Get token
			token := r.Header.Get("Authorization")
			if token == "" {
				helpers.SendError(w, nil, http.StatusUnauthorized, "Missing token")
				return
			}

			// Parse token
			claims, err := jwt.GetUserFromJwt(token)
			if err != nil {
				helpers.SendError(w, err, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Get userID
			userIDVal, ok := claims["userId"]
			if !ok {
				helpers.SendError(w, nil, http.StatusUnauthorized, "Unauthorized")
				return
			}

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

			// Fetch user
			user, err := userUC.GetUserById(userID)
			if err != nil || user == nil {
				helpers.SendError(w, errors.New("user not found"), http.StatusNotFound, "User not found")
				return
			}

			// Role check
			if !allowedRoles[user.Role] {
				helpers.SendError(w, nil, http.StatusForbidden, "Unauthorized: insufficient role")
				return
			}

			// Add user to context
			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
