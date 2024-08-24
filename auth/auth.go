package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecret"

func GenerateJWT(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
