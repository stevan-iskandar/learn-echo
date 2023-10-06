package controllers

import (
	"fmt"
	"learn-echo/helpers"
	"learn-echo/models"
	"learn-echo/structs"
	"net/http"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Root(c echo.Context) error {
	// Test write speed.
	startTime := time.Now()

	permissions := &[]models.Permission{}
	err := mgm.Coll(&models.Permission{}).SimpleFind(permissions, map[string]interface{}{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	var allowedPermissions []string
	for _, permission := range *permissions {
		allowedPermissions = append(allowedPermissions, permission.Code)
	}

	for i := 1; i <= 1000; i++ {
		password, _ := helpers.HashPassword(fmt.Sprintf("password*%d*", i))
		user := &models.User{
			Username:    fmt.Sprintf("user%d", i),
			Email:       fmt.Sprintf("user%d@email.com", i),
			Password:    password,
			FirstName:   fmt.Sprintf("first%d", i),
			LastName:    fmt.Sprintf("last%d", i),
			WrongPass:   i % 2,
			Permissions: allowedPermissions,
		}
		if err := helpers.FirstOrCreate(&models.User{}, bson.M{"username": user.Username}, user); err != nil {
			return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
		}
	}

	writeDuration := fmt.Sprintf("Write Time: %v\n", time.Since(startTime))

	// Test read speed.
	startTime = time.Now()
	var users []models.User
	if err := mgm.Coll(&models.User{}).SimpleFind(&users, bson.M{}); err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	readDuration := fmt.Sprintf("Read Time: %v\n", time.Since(startTime))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"write_duration": writeDuration,
		"read_duration":  readDuration,
		"users":          users,
	})
}
