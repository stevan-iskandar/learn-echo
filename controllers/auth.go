package controllers

import (
	"net/http"
	"time"

	"learn-echo/helpers"
	"learn-echo/middlewares"
	"learn-echo/validations"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c echo.Context) error {
	// Replace this with your registration logic, e.g., storing user data in the database.
	// After registration, you can generate and return a JWT to the client.

	// For simplicity, we'll generate a token with some mock user data.
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	token, err := helpers.GenerateToken("user123", "john.doe", expirationTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func getUserByUsername(username string) (*UserMGM, error) {
	filter := bson.M{"username": username}

	user := &UserMGM{}

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
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err,
		})
	}

	if validated := helpers.VerifyPassword(user.Password, credentials.Password); !validated {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Wrong password",
		})
	}

	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	token, err := helpers.GenerateToken(user.ID.String(), user.Username, expirationTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":      token,
		"user":       user,
		"expires_at": expirationTime,
	})
}

func GetUser(c echo.Context) error {
	claims := c.Get("user").(*middlewares.CustomClaims)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":  claims.UserID,
		"username": claims.Username,
	})
}
