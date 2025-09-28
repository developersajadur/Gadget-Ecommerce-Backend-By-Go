package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId int    `json:"userId"`
	Name  string `json:"name"`
	Age  int    `json:"age"`
	Email  string `json:"email"`

}

func GenerateJWT(secret []byte, payload JwtCustomClaims) (string, error) {
	claims := jwt.MapClaims{
		"userId": payload.UserId,
		"email":  payload.Email,
		"name":   payload.Name,
		"age":    payload.Age,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	// Use HS256 (not ES256) for byte slice secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

