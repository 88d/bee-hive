package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	UserID string   `json:"userID"`
	Scopes []string `json:"scopes"`
	jwt.StandardClaims
}

// HasAccess is used to check if this token has access to the provided access
func (j *JwtCustomClaims) HasAccess(scope string) bool {
	return isStringInArray(j.Scopes, scope)
}

func ClaimsFromRequestQuery(r *http.Request) (*JwtCustomClaims, error) {
	tokenString := r.URL.Query().Get("token")
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(config.SigningKey), nil
	})
	if err == nil && token.Valid {
		return token.Claims.(*JwtCustomClaims), nil
	}
	return nil, err
}
