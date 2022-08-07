// The business logic.
// versions:
// - protoc-gen-echo v0.0.1
// - protoc  v3.14.0
// source: example/v1/greeter.proto

package v1

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GreeterSayHelloBusinessHandler(req *HelloRequest) (HelloReply, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return HelloReply{}, err
	}
	fmt.Printf("Got HelloRequest is: %v\n", string(reqJson))
	return HelloReply{}, nil

}

func GreeterCreateUserBusinessHandler(req *CreateUserRequest) (CreateUserReply, error) {
	// Throws unauthorized error
	if req.Username != "hello" || req.Password != "world" {
		return CreateUserReply{}, echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		"Hello World",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwtKey := "dangerous"
	if os.Getenv("JWT-KEY") != "" {
		jwtKey = os.Getenv("JWT-KEY")
	}
	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return CreateUserReply{}, err
	}

	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return CreateUserReply{}, err
	}
	fmt.Printf("Got CreateUserRequest is: %v\n", string(reqJson))
	return CreateUserReply{Token: "Bearer " + t}, nil
}

func GreeterUpdateUserBusinessHandler(req *UpdateUserRequest) (UpdateUserReply, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return UpdateUserReply{}, err
	}
	fmt.Printf("Got UpdateUserRequest is: %v\n", string(reqJson))
	return UpdateUserReply{}, nil

}

func GreeterDeleteUserBusinessHandler(req *UserRequest) (UserReply, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return UserReply{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(reqJson))
	return UserReply{}, nil

}

func GreeterListUsersBusinessHandler(req *UserRequest) (UserReplys, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return UserReplys{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(reqJson))
	return UserReplys{}, nil

}
