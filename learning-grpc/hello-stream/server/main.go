package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
	hellowstream "tomgs-go/learning-grpc/hello-stream/api"
)

////
type server3 struct {
	hellowstream.UnimplementedStreamGreeterServer
}

func (s *server3) ResponseStream(r *hellowstream.StreamRequest, stream hellowstream.StreamGreeter_ResponseStreamServer) error {
	for n := 0; n <= 6; n++ {
		if n == 3 {
			return status.Error(10010001, "invalid argument")
		}
		err := stream.Send(&hellowstream.StreamReply{Message: "reply: " + r.Name + ":" + strconv.Itoa(n)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server3) ClientStream(stream hellowstream.StreamGreeter_ClientStreamServer) error  {
	return status.Error(10010001, "invalid argument")
}

func (s *server3) BidirectionalStream(stream hellowstream.StreamGreeter_BidirectionalStreamServer) error {
	return status.Error(10010001, "invalid argument")
}

// https://juejin.cn/post/6844903957794390024#heading-7
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	hellowstream.RegisterStreamGreeterServer(s, &server3{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
