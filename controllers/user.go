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

func UserList(c echo.Context) error {
	var users []*models.User
	usernames := c.QueryParams().Get("username")
	email := c.QueryParam("email")
	data, err := helpers.Pagination(c, &models.User{}, users, map[string]interface{}{
		"username": usernames,
		"email":    email,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, structs.Response{
		Message: constants.SM_RETRIEVE_SUCCESS,
		Data:    data,
	})
}

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

func UserUpdate(c echo.Context) error {
	strVal := c.QueryParam("id")

	encryptedData := helpers.Encrypt(strVal)
	decryptedData := helpers.Decrypt(encryptedData)

	return c.JSON(http.StatusOK, structs.Response{
		Data: map[string]interface{}{
			"string":         strVal,
			"encrypted_data": encryptedData,
			"decrypted_data": decryptedData,
		},
	})
}
