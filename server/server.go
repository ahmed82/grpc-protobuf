package main

import (
	"context"
	pb "grpc-protobuf/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {
	/// CREATE TCP SERVER
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	/// CREATE gRPC SERVER
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	/// Connect gRpc & TCP SERVERS
	err = grpcServer.Serve(listner)

	if err != nil {
		log.Println(err)
	}
}
