package main

import (
	"context"
	"flag"
	pb "learning-grpc/pb/src/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultMsg = "hello"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
	msg  = flag.String("msg", defaultMsg, "say hello")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf(*msg)
	r, err := c.PrintHelloWorld(ctx, &pb.HelloRequest{Message: *msg})
	if err != nil {
		log.Fatalf("error when sending request: %v", err)
	}

	log.Printf("Received: %s", r.GetMessage())
}
