package auth

import (
	"errors"
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

const (
	bearer = "Bearer"
	header = "Authorization"
)

// jwtFromHeader returns a `jwtExtractor` that extracts token from request header.
func JWTFromHeader(r *http.Request) (string, error) {
	auth := r.Header.Get(header)
	l := len(bearer)
	if len(auth) > l+1 && auth[:l] == bearer {
		return auth[l+1:], nil
	}
	return "", errors.New("empty or invalid jwt in request header")
}

func ClaimsFromRequest(r *http.Request) (*JwtCustomClaims, error) {
	tokenString, err := JWTFromHeader(r)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, config.Claims, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return config.SigningKey, nil
	})
	if err == nil && token.Valid {
		return token.Claims.(*JwtCustomClaims), nil
	}
	return nil, err
}
