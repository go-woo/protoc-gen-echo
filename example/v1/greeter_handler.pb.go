// The business logic.
// versions:
// - protoc-gen-echo v0.0.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

func GreeterSayHelloBusinessHandler(pathParam *map[string]string, payload *HelloRequest) (HelloReply, error) {
	// Here can put your business logic
	return HelloReply{}, nil
}

func GreeterCreateUserBusinessHandler(pathParam *map[string]string, payload *UserRequest) (UserReply, error) {
	// Here can put your business logic
	return UserReply{}, nil
}

func GreeterUpdateUserBusinessHandler(pathParam *map[string]string, payload *UserRequest) (UserReply, error) {
	// Here can put your business logic
	return UserReply{}, nil
}

func GreeterDeleteUserBusinessHandler(pathParam *map[string]string, payload *UserRequest) (UserReply, error) {
	// Here can put your business logic
	return UserReply{}, nil
}

func GreeterListUsersBusinessHandler(pathParam *map[string]string, payload *UserRequest) (UserReplys, error) {
	// Here can put your business logic
	return UserReplys{}, nil
}
