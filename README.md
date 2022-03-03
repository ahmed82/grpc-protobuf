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

NOTE: to avoid generating proto file by commend line, used Makefile.

# GO VSCode Autocomplete and Auto-import

1- Verify the GOPATH: *OPTINALY
Install golang as usual, and setting Gopath manually first. I make go-wokspace folder on C:\Users\{user}\Documents which is we can use as GOPATH. After that you add on user variables. Under go-workspace folder, create three folder called bin, pkg, and src

```go
go env
```
2- Open Visual Studio and search Golang Plugin and then install it
3- VsCode click View -> Command Pallete or type <b>Ctrl+Shift+P<b> and type "goinstall update/tools".
4- check all dependencies and click OK. it will take time to download all dependencies.
5- Add custom configuration on User Setting/ <b>Ctrl+Shift+P<b> and type "User" select "Preferences: Open User Settings"
on settings.json, Add this script and then restart vscode.
 
![image](https://user-images.githubusercontent.com/9446035/156212052-8a980390-fae6-44a7-bb2d-03e7560d36ce.png)


# package issues:
running "go mod tidy" solved the issue

This command goes through the go.mod file to resolve dependencies:

delete the packages that are not needed
download those needed
update the go.sum

# Run the gRPC Application

1- Build the project
```shell
make clean & make create
```
2- Run  the Server
```shell
cd server
go run .\server.go 
```
3- Run the Clint to test the API respons
```shell
cd .\client\
go run .\client.go
```
-----------------------------------------------------------------------------------------
# Creating REST API based on gRPC with Gateway
```
https://github.com/grpc-ecosystem/grpc-gateway
```
new file in the project root tools.go
```go
// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	
```
Run
```go
go mod tidy
```
