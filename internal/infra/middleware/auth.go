package middleware

// import (
// 	"context"
// 	"ecommerce/internal/domain"
// 	"ecommerce/internal/infra/repository"
// 	"ecommerce/pkg/helpers"
// 	"ecommerce/pkg/utils"
// 	"ecommerce/pkg/utils/jwt"
// 	"net/http"
// 	"strconv"
// )

// type contextKey string

// const UserContextKey = contextKey("user")

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		token := r.Header.Get("Authorization")
// 		if token == "" {
// 			helpers.SendError(w, nil, http.StatusUnauthorized, "Missing token")
// 			return
// 		}

// 		claims, err := jwt.GetUserFromJwt(token)
// 		if err != nil {
// 			helpers.SendError(w, err.Error(), http.StatusUnauthorized, "Invalid token")
// 			return
// 		}

// 		userID, ok := claims["userId"].(float64)
// 		if !ok {
// 			helpers.SendError(w, nil, http.StatusUnauthorized, "Unauthorized")
// 			return
// 		}
// 		if err != nil {
// 			helpers.SendError(w, err.Error(), http.StatusUnauthorized, "User not found")
// 			return
// 		}



// 		if user == nil {
// 			helpers.SendError(w, nil, http.StatusUnauthorized, "User not found")
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), UserContextKey, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
