package util

import (
	"errors"
	"fmt"
	"gatewayservice/internal/dto/resp"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateSignedJwt(userID int, userEmail string) (string, error) {
	const scope = "helper#GenerateSignedJwt"
	jwtExpiresAt, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_AT"))
	if err != nil {
		return "", fmt.Errorf("%s: %w", scope, errors.New("invalid environment variable"))
	}
	claims := &resp.JwtClaimsDto{
		ID:    userID,
		Email: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("%s: %w", scope, err)
	}
	return ss, nil
}
