package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "learning-grpc/pb/src/proto"
	"log"
	"time"
)

const (
	defaultMsg = "hello"
)

var (
	grpcTimeout    = flag.Duration("grpc_timeout", time.Second*2, "grpc timeout")
	grpcServerPort = flag.Int("grpcServerPort", 50051, "gRPC server port")
	addr           = flag.String("addr", "localhost", "server address")
	msg            = flag.String("msg", defaultMsg, "hello")
)

func main() {
	flag.Parse()
	serverAddr := fmt.Sprintf("%s:%d", *addr, *grpcServerPort)
	conn, err := grpc.NewClient(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("did not connect. err:%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("cause grpc client connection close error. err:%v", err)
		}
	}(conn)
	c := pb.NewHelloWorldClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), *grpcTimeout)
	defer cancel()
	log.Printf(*msg)
	r, err := c.PrintHelloWorld(ctx, &pb.HelloRequest{Message: *msg})
	if err != nil {
		log.Fatalf("error when sending request. err:%v", err)
	}

	log.Printf("Received message: %s", r.GetMessage())
}
