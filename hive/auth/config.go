package auth

type Config struct {
	SigningKey  string
	Claims      *JwtCustomClaims
	ExpireHours int
}
