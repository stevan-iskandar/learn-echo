package main

import (
	"fmt"
	_ "learn-echo/autoload"
	"learn-echo/constants"
	"learn-echo/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	routes.Init(e)

	// Run Server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv(constants.ENV_PORT))))
}
