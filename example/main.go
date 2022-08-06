package main

import (
	v1 "github.com/go-woo/protoc-gen-echo/example/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1.RegisterGreeterRouter(e)
	//you can add custom router outside protoc-gen-echo too.
	//MyCustomRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
