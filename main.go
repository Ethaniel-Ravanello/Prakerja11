package main

import (
	"keterampilan/config"
	routes "keterampilan/route"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDatabase()
	e := echo.New()
	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1123"))
}
