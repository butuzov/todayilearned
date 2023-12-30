package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	pb "github.com/butuzov/sandbox/grpc/server-streaming/proto"

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
		c   = pb.NewCalculatorClient(conn)
		in  = &wrapperspb.UInt64Value{Value: uint64(40)}
		ctx = context.Background()
	)

	if stream, err := c.Fibonacci(ctx, in); err != nil {
		log.Fatal(err)
	} else {
		done := make(chan bool)

		go func() {
			var i int
			for {
				resp, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					close(done)
					return
				}

				if err != nil {
					close(done)
					fmt.Errorf("cannot receive %v", err)
					return
				}

				log.Printf("Got@%d < %d \n", i, resp.GetValue())
				i++
			}
		}()

		<-done
	}
}
