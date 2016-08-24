package auth

import jwt "github.com/dgrijalva/jwt-go"

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	UserID string   `json:"userID"`
	Access []string `json:"access"`
	jwt.StandardClaims
}

// HasAccess is used to check if this token has access to the provided access
func (j *JwtCustomClaims) HasAccess(access string) bool {
	for _, a := range j.Access {
		if a == access {
			return true
		}
	}
	return false
}
