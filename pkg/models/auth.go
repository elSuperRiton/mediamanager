package models

import (
	"encoding/json"
)

type (
	// JWTLoginDTO is the data transfer object
	// for performing authentication using the JWT
	// strategy
	JWTLoginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// JWTLoginDAO is the data access object
	// returned by a succesfull JWT login operation
	JWTLoginDAO struct {
		Token string `json:"token"`
	}
)

// MarshalJSON allows for custom encoding of JWTLoginDAO in order
// to add context about the signing method
func (jwtDAO JWTLoginDAO) MarshalJSON() ([]byte, error) {
	type Alias JWTLoginDAO
	return json.Marshal(&struct {
		Type string `json:"type"`
		*Alias
	}{
		Alias: (*Alias)(&jwtDAO),
		Type:  "jwt",
	})
}
