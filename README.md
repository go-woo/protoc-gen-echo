## What is this?
This is a protocol plug-in that generates [echo](https://github.com/labstack/echo) server code 
from proto file.
## Installation
### Ubuntu
#### Pre installation
```
sudo apt install protobuf-compiler make
```

#### Generate example and run
```
make example
cd example && go run main.go
```
## Generate
```
$ protoc --proto_path=. \
    --proto_path=./third_party \
    --go_out=paths=source_relative:. \
    --echo_out=paths=source_relative:. \
    $(your_xxxx.proto)
```
[protoc](https://github.com/protocolbuffers/protobuf) help for using protocol.
## Write business logic
Your business logic stubs has been generated in `your_xxxx_handler.pb.go`, 
you can edit business logic in stubs.
```
func $(YourService)$(RpcName)BusinessHandler(pathParam *map[string]string, payload *YourRequest) 
    (YourReply, error) {
// Here can put logic
return YourReply{}, nil
}
```
All need func can be found in `your_xxxx_router.pb.go`.

### :bangbang: Caution

`your_xxxx_handler.pb.go` was generated only when it is generated for the first time, and will not be 
generated or overwritten after that, because the business logic code you added is already in it.

## Todo
- [ ] Query parameter
- [ ] Jwt and scope
- [ ] Casbin and scope
- [ ] error_reason