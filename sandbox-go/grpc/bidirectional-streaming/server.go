package main

import (
	"container/ring"
	"errors"
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/butuzov/sandbox/grpc/bidirectional-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:4772"
)

type server struct {
	pb.UnsafeCalculatorServer
}

func (s *server) AvgLast10(stream pb.Calculator_AvgLast10Server) error {
	// receive

	var (
		wg    sync.WaitGroup
		chRes = make(chan float64)
		cdSig = make(chan struct{})
		data  = ring.New(1)
	)
	wg.Add(3)

	stream.Send(wrapperspb.Double(1.0))

	// receiver
	go func() {
		defer wg.Done()
		defer func() { close(cdSig) }()

		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}

			if data.Len() < 10 {
				data.Link(&ring.Ring{Value: resp.GetValue()})
			} else {
				data.Value = resp.GetValue()
				data = data.Next()

			}
			// Adding value to ring
			cdSig <- struct{}{}

		}
	}()

	// worker
	go func() {
		defer wg.Done()

		for _ = range cdSig {
			var sum float64

			data.Do(func(i interface{}) {
				if value, ok := i.(uint64); ok {
					sum += float64(value)
				}
			})

			chRes <- (float64(sum) / float64(data.Len()))

		}
	}()

	// sender
	go func() {
		defer wg.Done()

		for value := range chRes {
			stream.Send(wrapperspb.Double(value))
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Printf("SERVER Runs @ %s \n", conn.Addr().String())

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)

	log.Fatal("failed to server: %v", s.Serve(conn))
}
