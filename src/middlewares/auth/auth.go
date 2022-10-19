package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJwt(userID uuid.UUID) (tokenString string, err error) {

	type JwtClaims struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Create the Claims
	claims := JwtClaims{
		"someone@example.com",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "vehicle-api",
			Subject:   userID.String(),
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtSecret)

	return
}
