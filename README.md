"# grpc-protobuf"

### In this tutorial, we build a basic API using gRPC and protobufs in Go

## Run `go run main.go` to run the app, run `go build main.go` to build an executable file.

## Download the protobuff
[]  https://developers.google.com/protocol-buffers/
[]  https://github.com/protocolbuffers/protobuf/releases

export the path
```ps
protoc --version
```
## Install library
```go
go get -u google.golang.org/grpc
```
```go
go get -u github.com/golang/protobuf/protoc-gen-go
The terminal replace the above commend with the nextone
go install github.com/golang/protobuf/protoc-gen-go@latest
## The website advice to use
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
go mod why github.com/golang/protobuf will show the full chain importing the package back to the direct dependency.


## Start new project

if you use golang create new. or from the cmd 
```go
go mod init [yourModule.name]
```
Now run the compiler, specifying the source directory (where your application's source code lives – the current directory is used if you don't provide a value), the destination directory (where you want the generated code to go; often the same as $SRC_DIR), and the path to your .proto. In this case, you would invo
```
protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto

protoc -I=$SRC_DIR --go_out=proto service.proto --proto_path=proto  --proto_path=third_party 

protoc --go_out=. --go_opt=path=source_relative --go-grpc_out=. --go-grpc_opt=path=source_relative proto/service.proto
```