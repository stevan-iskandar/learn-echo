package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/stevan-iskandar/learn-echo/controllers"
	"github.com/stevan-iskandar/learn-echo/validations"
)

func Init(e *echo.Echo) {
	e.GET("/", controllers.Root)

	e.POST("/user", controllers.Store, validations.StoreValidation)
}
