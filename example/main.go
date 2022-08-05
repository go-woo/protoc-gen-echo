package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "protoc-gen-echo/example/v1"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1.RegisterGreeterRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
