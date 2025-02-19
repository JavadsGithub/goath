package auth

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type Authenticator interface {
	authenticate(tokenString string) bool
}

func Authenticate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// TODO: make the secret dynamic
		return []byte("It'sMySecret"), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// TODO: evaluate claims and vlidate the token
		fmt.Println(claims)
		return true, nil
	} else {
		fmt.Println(err)
		return false, fmt.Errorf("Token is not valid")
	}
}
