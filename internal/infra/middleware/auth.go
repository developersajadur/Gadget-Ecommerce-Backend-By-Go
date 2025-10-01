package middleware

// import (
// 	"context"
// 	"ecommerce/pkg/helpers"
// 	"net/http"
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

// 		claims, err := utils.GetUserFromJwt(token)
// 		if err != nil {
// 			helpers.SendError(w, err.Error(), http.StatusUnauthorized, "Invalid token")
// 			return
// 		}

// 		userID, ok := claims["userId"].(float64)
// 		if !ok {
// 			helpers.SendError(w, nil, http.StatusUnauthorized, "Unauthorized")
// 			return
// 		}

// 		var user *database.Person
// 		for _, p := range database.People {
// 			if p.ID == int(userID) {
// 				user = &p
// 				break
// 			}
// 		}

// 		if user == nil {
// 			helpers.SendError(w, nil, http.StatusUnauthorized, "User not found")
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), UserContextKey, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
