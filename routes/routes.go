package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stevan-iskandar/learn-echo/controllers"
)

func Init(e *echo.Echo) {
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\nThis is about page\n")
	})

	e.POST("/user", controllers.Store)
}
