// The business logic.
// versions:
// - protoc-gen-echo v0.0.1
// - protoc  v3.14.0
// source: example/v1/greeter.proto

package v1

import (
	"encoding/json"
	"fmt"
)

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

func GreeterCreateUserBusinessHandler(req *UserRequest) (UserReply, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return UserReply{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(reqJson))

	return UserReply{}, nil
}

func GreeterUpdateUserBusinessHandler(req *UserRequest) (UserReply, error) {
	// Here can put your business logic,protoc-gen-ent soon coming
	//Below is example business logic code

	reqJson, err := json.Marshal(req)
	if err != nil {
		return UserReply{}, err
	}
	fmt.Printf("Got UserRequest is: %v\n", string(reqJson))

	return UserReply{}, nil
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
