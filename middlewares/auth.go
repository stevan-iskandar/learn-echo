package middlewares

import (
	"learn-echo/constants"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		claims := &CustomClaims{}

		// Remove "Bearer " prefix from token string
		tokenString = tokenString[len("Bearer "):]

		// Parse and verify the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv(constants.JWT_KEY)), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", claims)
		return next(c)
	}
}
