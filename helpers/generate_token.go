package helpers

import (
	"learn-echo/constants"
	"learn-echo/models"
	"learn-echo/structs"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *models.User, expirationTime time.Time) (string, error) {
	claims := structs.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:     user.ID.Hex(),
			Issuer: user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv(constants.ENV_JWT_KEY)))
}
