package main

import (
	"github.com/labstack/echo/v4"
	v1 "protoc-gen-echo/example/v1"
)

func main() {
	e := echo.New()

	v1.RegisterGreeterHTTPServer(e)

	e.Logger.Fatal(e.Start(":1323"))
}
