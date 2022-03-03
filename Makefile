create:
	protoc --proto_path=proto proto/*.proto --go_out=gen
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen

clean: 
	rm gen/proto/*.go

## for the Getway 
#mkdir -p google/api
#curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
#curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto
