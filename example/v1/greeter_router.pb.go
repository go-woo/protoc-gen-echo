// Code generated by protoc-gen-echo. DO NOT EDIT.
// versions:
// - protoc-gen-echo v0.0.1
// - protoc  v3.14.0
// source: example/v1/greeter.proto

package v1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RegisterGreeterRouter(e *echo.Echo) {
	e.GET("/helloworld/:name/hi/:nice", _Greeter_SayHello0_HTTP_Handler)
	e.POST("/usr/:name/pwd/:passwd", _Greeter_CreateUser0_HTTP_Handler)
	e.PATCH("/usr/:name/pwd/:passwd", _Greeter_UpdateUser0_HTTP_Handler)
	e.DELETE("/usr/:name/pwd/:passwd", _Greeter_DeleteUser0_HTTP_Handler)
	e.GET("/usr/:name/pwd/:passwd", _Greeter_ListUsers0_HTTP_Handler)
}

func _Greeter_SayHello0_HTTP_Handler(c echo.Context) error {
	var req *HelloRequest = new(HelloRequest)
	req.Name = c.Param(strings.ToLower("Name"))
	req.Nice = c.Param(strings.ToLower("Nice"))

	reply, err := GreeterSayHelloBusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}

func _Greeter_CreateUser0_HTTP_Handler(c echo.Context) error {
	var req *UserRequest = new(UserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	reply, err := GreeterCreateUserBusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}

func _Greeter_UpdateUser0_HTTP_Handler(c echo.Context) error {
	var req *UserRequest = new(UserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	reply, err := GreeterUpdateUserBusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}

func _Greeter_DeleteUser0_HTTP_Handler(c echo.Context) error {
	var req *UserRequest = new(UserRequest)
	req.Phone = c.Param(strings.ToLower("Phone"))
	req.Email = c.Param(strings.ToLower("Email"))

	reply, err := GreeterDeleteUserBusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}

func _Greeter_ListUsers0_HTTP_Handler(c echo.Context) error {
	var req *UserRequest = new(UserRequest)
	req.Phone = c.Param(strings.ToLower("Phone"))
	req.Email = c.Param(strings.ToLower("Email"))

	reply, err := GreeterListUsersBusinessHandler(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &reply)
}
