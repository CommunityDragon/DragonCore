package main

import (
	"dragonback/modules"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

func main() {
	e := echoswagger.New(echo.New(), "/swagger", nil)
	e.
	modules.Bundler.Register(e)
	e.Echo().Logger.Fatal(e.Echo().Start(":80"))
}
