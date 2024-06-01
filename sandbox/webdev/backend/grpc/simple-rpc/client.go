package main

import (
	"context"
	"log"

	pb "github.com/butuzov/sandbox/grpc/simple-rpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:4772"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewServiceClient(conn)

	if response, err := c.Hello(context.Background(), &pb.Request{Message: "Ping"}); err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	} else {
		log.Printf("(incoming client) %#v \n", response.GetMessage())
	}
}
