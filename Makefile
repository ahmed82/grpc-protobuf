create:
	protoc --proto_path=proto proto/*.proto --out_out=gen/proto
	protoc --proto_path=proto proto/*.proto --out_grpc_out=gen/proto

clean: gen/proto/*.go