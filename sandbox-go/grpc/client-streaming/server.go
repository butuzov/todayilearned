package main

import (
	"errors"
	"io"
	"log"
	"net"

	pb "github.com/butuzov/sandbox/grpc/client-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:4772"
)

// gRPC Server
var _ pb.CalculatorServer = (*server)(nil)

type server struct {
	pb.UnsafeCalculatorServer
}

func (s *server) Avg(stream pb.Calculator_AvgServer) error {
	var (
		sum uint64
		cnt int
	)

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&wrapperspb.DoubleValue{
				Value: float64(sum) / float64(cnt),
			})
		}

		log.Printf("server got: %v", res.GetValue())
		sum = sum + res.GetValue()
		cnt++

	}

	return nil
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
