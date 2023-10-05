package routes

import (
	"learn-echo/controllers"
	"learn-echo/middlewares"
	"learn-echo/validations"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/", controllers.Root)
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login, validations.LoginValidation)

	api.Use(middlewares.Auth)

	RouteUser(api)
}
