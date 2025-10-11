package middleware

import (
	"context"
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"ecommerce/pkg/utils/jwt"
	"errors"
	"net/http"
	"time"
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

			jwtUser, err := jwt.GetUserFromJwt(r)
			if err != nil {
				helpers.SendError(w, err, http.StatusUnauthorized, "Invalid token")
				return
			}

			if time.Now().Unix() > jwtUser.Exp {
				helpers.SendError(w, errors.New("token expired"), http.StatusUnauthorized, "Token has expired")
				return
			}

			user, err := userUC.GetUserById(jwtUser.UserID)
			if err != nil || user == nil {
				helpers.SendError(w, errors.New("user not found"), http.StatusNotFound, "User not found")
				return
			}

			if user.IsBlocked {
				helpers.SendError(w, errors.New("user is blocked"), http.StatusForbidden, "User is blocked")
				return
			}

			if !user.IsVerified {
				helpers.SendError(w, errors.New("user is not verified"), http.StatusForbidden, "User is not verified")
				return
			}

			if user.IsDeleted {
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
