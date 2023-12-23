package main

import (
	"context"
	"flag"
	"fmt"
	pb "learning-grpc/pb/src/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "server port")
)

// HelloWorldServerを実装
type HelloWorldServer struct{}

func (hws *HelloWorldServer) PrintHelloWorld(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", hr.GetMessage())
	return &pb.HelloResponse{Message: hr.GetMessage() + "World"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &HelloWorldServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
