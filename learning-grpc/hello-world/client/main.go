package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	helloworldpb "tomgs-go/learning-grpc/hello-world/api"
)

func main() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Register Greeter
	client := helloworldpb.NewGreeterClient(conn)

	SayGood(err, client)
	SayHello(err, client)
}

func SayGood(err error, client helloworldpb.GreeterClient) {
	good, err := client.SayGood(context.Background(), &helloworldpb.HelloRequest{
		Name: "hello",
	})
	if err != nil {
		log.Println("Failed to invoke SayGood:", err)
	}

	fmt.Println(good)
}

func SayHello(err error, client helloworldpb.GreeterClient) {
	hello, err := client.SayHello(context.Background(), &helloworldpb.HelloRequest{
		Name: "hello",
	})
	if err != nil {
		log.Println("Failed to invoke SayHello:", err)
	}

	fmt.Println(hello)
}
