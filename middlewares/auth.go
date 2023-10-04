package middlewares

import (
	"learn-echo/constants"
	"learn-echo/structs"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const USER = "userAuth"

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{Message: "Missing token"})
		}

		// Remove "Bearer " prefix from token string
		tokenString = tokenString[len("Bearer "):]

		claims := &structs.JWTClaims{}

		// Parse and verify the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv(constants.ENV_JWT_KEY)), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{Message: err.Error()})
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, structs.Response{Message: "Invalid token"})
		}

		c.Set(USER, claims)
		return next(c)
	}
}
