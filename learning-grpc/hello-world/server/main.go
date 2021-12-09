package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"net"
	helloworldpb "tomgs-go/learning-grpc/hello-world/api"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	//return nil, status.Error(codes.Unimplemented, "{\"description\":\"未登录或登录信息过期\",\"errcode\":30000401}")
	return nil, status.Error(10010001, "invalid argument")
	//return nil, errors.New("test error")
	//return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func (s *server) SayGood(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error)  {
	return &helloworldpb.HelloReply{Message: in.Name + " Good"}, nil
}

////
type server2 struct {
	helloworldpb.UnimplementedGreeter2Server
}

func NewServer2() *server2 {
	return &server2{}
}

/*func (s *server2) SayHello2(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world2"}, nil
}*/

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
