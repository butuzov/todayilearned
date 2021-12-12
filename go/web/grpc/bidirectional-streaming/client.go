package main

import (
	"context"
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"

	pb "github.com/butuzov/sandbox/grpc/bidirectional-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:4772"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not connect with server %v", err)
	}
	defer conn.Close()

	var (
		client = pb.NewCalculatorClient(conn)
		ctx    = context.Background()
	)

	if stream, err := client.AvgLast10(ctx); err != nil {
		log.Fatal(err)
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		// Sender
		go func(stream pb.Calculator_AvgLast10Client) {
			defer func() { wg.Done() }()

			/*
				Sends some (19) random uint64 to the server
			*/

			for i := uint64(0); i < 19; i++ {
				value := uint64(rand.Int63n(599))
				if err := stream.Send(wrapperspb.UInt64(value)); err != nil {
					log.Fatalf("client_err: %v", err)
				} else {
					log.Printf("sent: %v", value)
				}
			}

			if err := stream.CloseSend(); err != nil {
				log.Fatal("CloseSend: %v", err)
			}
		}(stream)

		// Receiver
		go func(stream pb.Calculator_AvgLast10Client) {
			defer func() { wg.Done() }()

			for {
				resp, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					return
				}

				if err != nil {
				}

				log.Printf("client: avg %+v", resp.GetValue())

			}
		}(stream)

		wg.Wait()
	}
}
