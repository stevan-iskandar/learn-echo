package controllers

import (
	"net/http"

	"learn-echo/helpers"
	"learn-echo/middlewares"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	// Replace this with your registration logic, e.g., storing user data in the database.
	// After registration, you can generate and return a JWT to the client.

	// For simplicity, we'll generate a token with some mock user data.
	token, err := helpers.GenerateToken("user123", "john.doe")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func Login(c echo.Context) error {
	// Replace this with your login logic, e.g., verifying credentials.
	// After successful login, you can generate and return a JWT to the client.

	// For simplicity, we'll generate a token with some mock user data.
	token, err := helpers.GenerateToken("user123", "john.doe")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.CustomClaims)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":  claims.UserID,
		"username": claims.Username,
	})
}
