## 这是什么？
这是个protoc插件，通过定义proto文件生成[Echo](https://github.com/labstack/echo)服务端代码。

## 安装
### Ubuntu
#### 预安装
```
sudo apt install protobuf-compiler make
```
#### 运行例子
```
make example
cd example && go run main.go
```

## 生成
```
$   protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:. \
	       --echo_out=paths=source_relative:. \
	       $(your_xxxx.proto)
```
[protoc](https://github.com/protocolbuffers/protobuf)有protoc的使用帮助。
## 业务逻辑
你的业务逻辑桩已经生成在`your_xxxx_handler.pb.go`里，你可以在这些文件添加业务逻辑。
```
func $(YourService)$(RpcName)BusinessHandler(pathParam *map[string]string, payload *YourRequest) (YourReply, error) {
	// Here can put logic
	return YourReply{}, nil
}
```
所有的func都可以在对应的`your_xxxx_router.pb.go`里找到。
### :bangbang:注意
`your_xxxx_handler.pb.go`只有第一次生成的时候才被生成，之后不再被生成，也不会被覆盖，因为里面已经有了你添加的业务逻辑代码。
## Todo
- [ ] Query parameter
- [ ] Jwt and scope
- [ ] Casbin and scope
- [ ] error_reason