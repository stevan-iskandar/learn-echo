package helpers

import (
	"learn-echo/middlewares"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID, username string) (string, error) {
	// Create a new JWT token with custom claims.
	claims := middlewares.CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middlewares.SIGNING_KEY))
}
