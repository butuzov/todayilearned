package main

import (
	"context"
	"log"
	"math/rand"

	pb "github.com/butuzov/sandbox/grpc/client-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:4772"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()

	var (
		client = pb.NewCalculatorClient(conn)
		ctx    = context.Background()
	)

	if stream, err := client.Avg(ctx); err != nil {
		log.Fatal(err)
	} else {
		for i := uint64(0); i < 10; i++ {
			n := &wrapperspb.UInt64Value{Value: uint64(rand.Int63n(599))}
			if err := stream.Send(n); err != nil {
				log.Fatalf("Send: %v", err)
			}
		}

		if resp, err := stream.CloseAndRecv(); err != nil {
			log.Fatalf("client/Error/CloseAndRecv: %v\n", err)
		} else {
			log.Printf("client/Result/CloseAndRecv: %v\n", resp.GetValue())
		}

	}
}
