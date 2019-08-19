package utils

import (
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestGenerateJWTToken(t *testing.T) {
	testSuite := []struct {
		secretKey []byte
		signedJWT string
		claims    jwt.MapClaims
	}{
		{
			secretKey: []byte("testsecretkey"),
			signedJWT: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDYzMDQ1MjEsIm5iZiI6MTU0NjMwNDUyMn0.37HReu_1mfbiUedFHZNwt9dpPRSbmQisu6QN6r0Ze4k",
			claims: jwt.MapClaims{
				"exp": time.Date(2019, 01, 01, 01, 01, 01, 01, time.UTC).Add(1 * time.Minute).Unix(),
				"nbf": time.Date(2019, 01, 01, 01, 01, 02, 02, time.UTC).Add(1 * time.Minute).Unix(),
			},
		},
		{
			secretKey: []byte("othertestsecrettoken"),
			signedJWT: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDYzMDQ1MjEsIm5iZiI6MTU0NjMwNDUyMn0.sYCEtVb18Hl7co65bOFrM3YyEHTkbRg_0cLXPuiWbMQ",
			claims: jwt.MapClaims{
				"exp": time.Date(2019, 01, 01, 01, 01, 01, 01, time.UTC).Add(1 * time.Minute).Unix(),
				"nbf": time.Date(2019, 01, 01, 01, 01, 02, 02, time.UTC).Add(1 * time.Minute).Unix(),
			},
		},
	}

	for _, test := range testSuite {
		token, err := GenerateJWTToken(test.secretKey, test.claims)
		if err != nil {
			t.Errorf("error generating token : %v", err)
		}

		if token != test.signedJWT {
			t.Errorf("expected token to equal %v, got %v", test.signedJWT, token)
		}
	}
}
