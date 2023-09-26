package controllers

import (
	"net/http"
	"time"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

// UserForm struct
type UserForm struct {
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Age       interface{} `json:"age"`
	Safe      interface{} `json:"safe"`
	Code      string      `json:"code"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func Store(c echo.Context) error {
	data, err := validate.FromRequest(c.Request())
	if err != nil {
		return err
	}

	v := data.Create()

	v.StringRules(map[string]string{
		"name":  "required|alphaDash|min_len:7",
		"email": "required|email",
		"age":   "required|min:13",
		"code":  `required|regex:\d{4,6}`,
		"safe":  "-",
	})

	if !v.Validate() {
		return c.JSON(http.StatusUnprocessableEntity, v.Errors)
	}

	userForm := &UserForm{}
	v.BindStruct(userForm)

	userForm.CreatedAt = time.Now()
	userForm.UpdatedAt = time.Now()

	return c.JSON(http.StatusOK, userForm)
}
