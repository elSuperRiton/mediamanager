package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken is a helper method for generating a JWT token signed with the
// provided secret and MapClaims
func GenerateJWTToken(secretKey []byte, mapClaims jwt.MapClaims) (signedToken string, signingErr error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	signedToken, signingErr = token.SignedString(secretKey)
	return
}
