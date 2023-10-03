package controllers

import (
	"net/http"

	"learn-echo/constants"
	"learn-echo/helpers"
	"learn-echo/models"
	"learn-echo/structs"
	"learn-echo/validations"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
)

func UserStore(c echo.Context) error {
	userForm := c.Get(validations.STORE_VALIDATION).(*validations.UserStoreForm)

	user := &models.User{
		Username:  userForm.Username,
		Email:     userForm.Email,
		FirstName: userForm.FirstName,
		LastName:  userForm.LastName,
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: constants.SM_STORE_SUCCESS,
		Data:    user,
	})
}

func UserList(c echo.Context) error {
	var users []*models.User
	data, err := helpers.Pagination(c, &models.User{}, users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: constants.SM_RETRIEVE_SUCCESS,
		Data:    data,
	})
}
