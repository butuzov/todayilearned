/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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
