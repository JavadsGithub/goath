package utils

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Authenticator interface {
	authenticate(tokenString string) bool
}

type ClaimsValidator func(claims jwt.MapClaims) error

func Authenticate(tokenString string, validator ClaimsValidator) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")
		if len(secret) < 1 {
			return false, errors.New("there is no secret")
		}
		// TODO: make the secret dynamic
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if err := validator(claims); err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		return false, errors.New("token is not valid")
	}
}
