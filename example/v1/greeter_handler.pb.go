// The business logic.
// versions:
// - protoc-gen-echo v0.1.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GreeterSayHelloBusinessHandler(req *HelloRequest, c echo.Context) (HelloReply, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	rj, err := json.Marshal(req)
	if err != nil {
		return HelloReply{}, err
	}
	fmt.Printf("Got HelloRequest is: %v\n", string(rj))
	return HelloReply{}, nil
}

func GreeterCreateUserBusinessHandler(req *CreateUserRequest, c echo.Context) (CreateUserReply, error) {
	// Throws unauthorized error
	if req.Username != "hello" || req.Password != "world" {
		return CreateUserReply{}, echo.ErrUnauthorized
	}
	// Set custom claims
	claims := &JwtCustomClaims{
		Name:  "Hello World",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	jk := "dangerous"
	if os.Getenv("JWTKEY") != "" {
		jk = os.Getenv("JWTKEY")
	}
	t, err := token.SignedString([]byte(jk))
	if err != nil {
		return CreateUserReply{}, err
	}

	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	rj, err := json.Marshal(req)
	if err != nil {
		return CreateUserReply{}, err
	}
	fmt.Printf("Got CreateUserRequest is: %v\n", string(rj))
	return CreateUserReply{Token: "Bearer " + t}, nil
}

func GreeterUpdateUserBusinessHandler(req *UpdateUserRequest, c echo.Context) (UpdateUserReply, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	username := claims.Name
	fmt.Printf("Got jwt name is: %v\n", username)
	req.Username = username

	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	rj, err := json.Marshal(req)
	if err != nil {
		return UpdateUserReply{}, err
	}
	fmt.Printf("Got UpdateUserRequest is: %v\n", string(rj))
	return UpdateUserReply{}, nil
}

func GreeterDeleteUserBusinessHandler(req *UserRequest, c echo.Context) (UserReply, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	rj, err := json.Marshal(req)
	if err != nil {
		return UserReply{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(rj))
	return UserReply{}, nil
}

func GreeterListUsersBusinessHandler(req *UserRequest, c echo.Context) (UserReplys, error) {
	// Here can put your business logic, can use ORM:github.com/go-woo/protoc-gen-ent
	// Below is example business logic code
	rj, err := json.Marshal(req)
	if err != nil {
		return UserReplys{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(rj))
	return UserReplys{}, nil
}
