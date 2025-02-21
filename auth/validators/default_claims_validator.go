package validators

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func DefaultClaimsValidator(claims jwt.MapClaims) bool {
	// Check if the token is expired
	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return false
		}
	}

	// // Check if the issuer is correct
	// if iss, ok := claims["iss"].(string); ok && iss != "your-issuer" {
	// 	return false
	// }

	// // Check if the audience is correct
	// if aud, ok := claims["aud"].(string); ok && aud != "your-audience" {
	// 	return false
	// }

	return true
}
