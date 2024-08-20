package auth

import (
	"time"

	"danieljonguitud.com/aws-events-go/db"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateJWT(user db.CreateUserRow) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
