package middlewares

import (
	"learn-echo/helpers"
	"learn-echo/models"
	"learn-echo/structs"
	"net/http"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Permission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userAuth := c.Get(USER).(*structs.JWTClaims)
			objectID, err := primitive.ObjectIDFromHex(userAuth.ID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
			}

			user := &models.User{}
			err = mgm.Coll(user).FindByID(objectID, user)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
			}

			if !helpers.StringExistsInArray(user.Permissions, permission) {
				return c.JSON(http.StatusUnauthorized, structs.Response{
					Message: "Unauthorized",
				})
			}
			return next(c)
		}
	}
}
