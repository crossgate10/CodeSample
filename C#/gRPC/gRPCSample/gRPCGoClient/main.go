package main

import (
	"log"
	"os"
	"time"

	pb "../gRPCDemo/protos"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:9007"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGRPCClient(conn)

	var name = "Ben"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Message: ", r.Message)

}
