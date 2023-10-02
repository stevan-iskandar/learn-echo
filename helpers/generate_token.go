package helpers

import (
	"learn-echo/constants"
	"learn-echo/middlewares"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, username string, expirationTime time.Time) (string, error) {
	claims := middlewares.CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(constants.JWT_KEY)))
}
