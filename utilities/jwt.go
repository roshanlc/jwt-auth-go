package utilities

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Signging Key
var SigningKey []byte

func GenerateJWT(email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, err
}
