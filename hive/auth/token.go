package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken string to use as needed
func GenerateToken(userID string, access []string) string {
	claims := &JwtCustomClaims{
		userID,
		access,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.ExpireHours)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		return ""
	}
	return t
}
