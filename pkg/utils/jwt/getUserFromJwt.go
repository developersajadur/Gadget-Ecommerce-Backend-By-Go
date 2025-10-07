package jwt

import (
	"ecommerce/internal/config"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string
	Email  string
	Role   string
	Exp    int64
}

func GetUserFromJwt(r *http.Request) (*UserClaims, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("missing token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ENV.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := ""
		switch v := claims["userId"].(type) {
		case float64:
			userID = fmt.Sprintf("%v", int64(v))
		case string:
			userID = v
		default:
			return nil, errors.New("invalid userId in token")
		}

		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)
		exp, _ := claims["exp"].(float64)

		return &UserClaims{
			UserID: userID,
			Email:  email,
			Role:   role,
			Exp:    int64(exp),
		}, nil
	}

	return nil, errors.New("invalid token")
}
