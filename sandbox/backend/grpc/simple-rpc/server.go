// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/butuzov/sandbox/grpc/simple-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:4772"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedServiceServer
}

var _ pb.ServiceServer = (*server)(nil)

// SayHello implements helloworld.GreeterServer
func (s *server) Hello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("(incoming server) %#v \n", in.GetMessage())

	switch in.GetMessage() {
	case "Ping":
		return &pb.Reply{Message: "Pong"}, nil
	}

	return &pb.Reply{Message: "<echo>" + in.GetMessage()}, nil
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("SERVER RUNS@", listener.Addr().String())

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})

	reflection.Register(s)

	log.Fatalf("failed to serve: %v", s.Serve(listener))
}
