# proto-practice

### 项目说明

单一的 protobuf 声明项目，方便 服务方 / 使用方 以 git submodule 的模式引用

### 项目地址

https://github.com/meizikeai/proto-practice.git

### 使用说明

执行 build.proto.sh 文件命令，依赖 Go 插件，如下

```go
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

插件更多资讯见 https://grpc.io/docs/languages/go/quickstart
