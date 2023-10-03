package controllers

import (
	"net/http"
	"time"

	"learn-echo/helpers"
	"learn-echo/middlewares"
	"learn-echo/models"
	"learn-echo/structs"
	"learn-echo/validations"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c echo.Context) error {
	return c.JSON(http.StatusOK, structs.Response{
		Message: "User created",
	})
}

func getUserByUsername(username string) (*models.User, error) {
	filter := bson.M{"username": username}

	user := &models.User{}

	// Find the user document in the collection
	err := mgm.Coll(user).First(filter, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(c echo.Context) error {
	credentials := c.Get(validations.LOGIN_VALIDATION).(*validations.Credentials)

	user, err := getUserByUsername(credentials.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, structs.Response{
			Message: err.Error(),
		})
	}

	if validated := helpers.VerifyPassword(user.Password, credentials.Password); !validated {
		return c.JSON(http.StatusUnauthorized, structs.Response{
			Message: "Wrong password",
		})
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	token, err := helpers.GenerateToken(user.ID, user.Username, expirationTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: "Successfully logged in",
		Data: map[string]interface{}{
			"token":      token,
			"user":       user,
			"expires_at": expirationTime,
		},
	})
}

func User(c echo.Context) interface{} {
	return c.Get(middlewares.USER).(*middlewares.CustomClaims)
}
