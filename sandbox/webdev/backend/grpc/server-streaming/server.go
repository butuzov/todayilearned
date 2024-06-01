// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	pb "github.com/butuzov/sandbox/grpc/server-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:4772"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnsafeCalculatorServer
}

var _ pb.CalculatorServer = (*server)(nil)

// SayHello implements helloworld.GreeterServer
func (s *server) Fibonacci(in *wrapperspb.UInt64Value, service pb.Calculator_FibonacciServer) error {
	log.Printf("We asked to generate first %d members of fibonacci sequence:\n", in.GetValue())

	ch := fibonacciSwaps(in.GetValue())

	for n := range ch {
		service.Send(&wrapperspb.UInt64Value{Value: n})
	}

	return nil
}

func fibonacciSwaps(n uint64) <-chan uint64 {
	out := make(chan uint64)

	go func() {
		var (
			a uint64 = 0
			b uint64 = 1
		)
		for i := uint64(0); i < n; i++ {
			out <- a
			a, b = b, a+b
		}
	}()

	return out
}

func main() {
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer conn.Close()
	log.Printf("SERVER Runs @ %s \n", conn.Addr().String())

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)

	log.Fatalf("failed to serve: %v", s.Serve(conn))
}
