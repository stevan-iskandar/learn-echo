package middlewares

import (
	"learn-echo/helpers"
	"learn-echo/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Permission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// user := c.Get(USER).(*structs.JWTClaims)

			if !helpers.StringExistsInArray([]string{}, permission) {
				return c.JSON(http.StatusUnauthorized, structs.Response{
					Message: "Unauthorized",
				})
			}
			return next(c)
		}
	}
}
