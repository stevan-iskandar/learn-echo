package routes

import (
	"learn-echo/controllers"
	"learn-echo/middlewares"
	"learn-echo/validations"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/", controllers.Root)

	e.POST("/user", controllers.Store, validations.StoreValidation)

	api := e.Group("/api")

	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	api.Use(middlewares.Auth)

	api.GET("/user", controllers.GetUser)
}
