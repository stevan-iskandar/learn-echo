package main

import (
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/stevan-iskandar/learn-echo/autoload"
	"github.com/stevan-iskandar/learn-echo/routes"
)

func main() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	routes.Init(e)

	// Run Server
	e.Logger.Fatal(e.Start(":8080"))
}
