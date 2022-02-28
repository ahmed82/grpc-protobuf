package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-protobuf/gen/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Clint test gRPC API response"})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	//or return the message only
	fmt.Println(resp.Msg)
}
