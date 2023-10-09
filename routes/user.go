package routes

import (
	"learn-echo/constants"
	"learn-echo/controllers"
	"learn-echo/middlewares"
	"learn-echo/validations"

	"github.com/labstack/echo/v4"
)

func RouteUser(api *echo.Group) {
	api.GET(
		"/user",
		controllers.UserList,
		middlewares.Permission(constants.PER_USER_VIEW),
	)
	api.POST(
		"/user",
		controllers.UserStore,
		middlewares.Permission(constants.PER_USER_CREATE),
		validations.StoreValidation,
	)
	api.GET(
		"/user/id",
		controllers.UserUpdate,
	)
}
