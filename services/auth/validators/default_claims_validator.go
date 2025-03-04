package validators

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func DefaultClaimsValidator(claims jwt.MapClaims) error {
	// Check if the token is expired
	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return errors.New("token expired")
		}
	}

	// TODO: Check if the issuer is correct
	if _, ok := claims["iss"].(string); ok {
		return errors.New("invalid issuer")
	}

	// TODO: Check if the audience is correct
	if _, ok := claims["aud"].(string); ok {
		return errors.New("invalid audience")
	}

	return nil
}
