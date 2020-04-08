package main

import (
	"dragoncore/lib/models/validator"
	"dragoncore/modules"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
)

func main() {
	e := echoswagger.New(echo.New(), "/swagger", nil)
	e.Echo().Validator = validator.New()
	e.Echo().Use(middleware.Logger())
	modules.Bundler.Register(e)

	e.Echo().Logger.Fatal(e.Echo().Start(":80"))
}
