package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stevan-iskandar/learn-echo/models"
	"github.com/stevan-iskandar/learn-echo/validations"
)

func Store(c echo.Context) error {
	userForm := c.Get(validations.STORE_VALIDATION).(*validations.UserStoreForm)

	user := models.User{
		Name:      userForm.Name,
		Email:     userForm.Email,
		Age:       userForm.Age,
		Safe:      userForm.Safe,
		Code:      userForm.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	usersColl := models.DB().Collection(models.USER)

	_, err := usersColl.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, user)
}
