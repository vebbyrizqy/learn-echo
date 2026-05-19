// @title Learn Echo API
// @version 1.0
// @description API documentation for Learn Echo
// @host localhost:1323
// @BasePath /
package main

import (
	"learn-echo/routes"
	"learn-echo/config"
	_ "learn-echo/docs"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

)

func main() {
	config.ConnectDatabase()

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}