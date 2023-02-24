package resp

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
