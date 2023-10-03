package validations

import (
	"net/http"

	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

// UserStoreForm struct
type UserStoreForm struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

const STORE_VALIDATION = "StoreValidation"

func StoreValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := validate.FromRequest(c.Request())
		if err != nil {
			return err
		}

		v := data.Create()

		v.FilterRules(map[string]string{
			"age":  "int",
			"safe": "int",
		})
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

		userForm := &UserStoreForm{}
		v.BindSafeData(userForm)
		c.Set(STORE_VALIDATION, userForm)

		return next(c)
	}
}
