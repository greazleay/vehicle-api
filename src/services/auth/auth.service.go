package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwt(userID uuid.UUID) (tokenString string, err error) {

	// Create the Claims
	claims := JwtClaims{
		"someone@example.com", // This value specified on puepose, to be changed later
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "vehicle-api",
			Subject:   userID.String(),
			Audience:  []string{"vehicle-api"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtSecret)

	return
}

func ValidateToken(signedToken string) (err error) {

	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Email, claims.RegisteredClaims.Issuer)
		return nil
	} else {
		err = errors.New("token expired")
		return
	}

}
