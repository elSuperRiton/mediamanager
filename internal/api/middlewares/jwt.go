package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elSuperRiton/mediamanager/pkg/utils"
)

type (
	JWTConfig struct {
		SigningKey []byte // Secret signing key
		HeaderKey  []byte // Name of the key ( ex: []byte{"Authorization"} )
	}
	jwtMdl struct {
		signingKey []byte
		headerKey  []byte
	}
)

// NewJWT returns a
func NewJWT(options *JWTConfig) MiddlewareFunc {
	j := &jwtMdl{
		signingKey: options.SigningKey,
		headerKey:  options.HeaderKey,
	}

	return j.Handler
}

// hasJWTToken is a helper method to retrieve JWT token in header of a request
func (j *jwtMdl) hasJWTToken(r *http.Request) (token string, hasToken bool) {
	if token := r.Header.Get(string(j.headerKey)); token != "" {
		return token, true
	}

	return "", false
}

// isValidJWT is a helper method for parsing token and verifying it's validity
// it uses "github.com/dgrijalva/jwt-go" under to hood to perform most tasks
func (j *jwtMdl) isValidJWT(jwtToken string) (isValidToken bool, parsingError error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return j.signingKey, nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	return false, nil
}

func (j *jwtMdl) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, hasToken := j.hasJWTToken(r)
		if !hasToken {
			utils.RenderErr(
				w,
				r,
				"No token provided in request",
				http.StatusUnauthorized,
			)

			return
		}

		isValid, err := j.isValidJWT(token)
		if err != nil {
			utils.RenderErr(
				w,
				r,
				"error validating token",
				http.StatusUnauthorized,
			)
			return
		}

		if !isValid {
			utils.RenderErr(
				w,
				r,
				"The provided token is invalid",
				http.StatusUnauthorized,
			)

			return
		}

		next.ServeHTTP(w, r)
	})
}
