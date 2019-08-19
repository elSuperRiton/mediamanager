package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/utils"
)

// Authentify is the handler responsible for performing authentication
func Authentify(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var loginDto models.JWTLoginDTO
	if err := decoder.Decode(&loginDto); err != nil {
		utils.RenderErr(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: replace with db checking
	if loginDto.Email != "test@test.com" || loginDto.Password != "pwd" {
		utils.RenderErr(w, r, "wrong email or password", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWTToken([]byte("mysupersecret"), jwt.MapClaims{})
	dao := models.JWTLoginDAO{
		Token: token,
	}

	utils.RenderData(w, r, dao, http.StatusOK)
	return
}
