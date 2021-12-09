package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	hellowstream "tomgs-go/learning-grpc/hello-stream/api"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := hellowstream.NewStreamGreeterClient(conn)

	err = responseStreamRequest(client, &hellowstream.StreamRequest{Name: "responseHello"})
	if err != nil {
		log.Fatalf("responseStreamRequest.err: %v", err)
	}

	err = clientStreamRequest(client, &hellowstream.StreamRequest{Name: "clientHello"})
	if err != nil {
		log.Fatalf("clientStreamRequest.err: %v", err)
	}

	err = bidirectionalStreamRequest(client, &hellowstream.StreamRequest{Name: "bidirectionalHello"})
	if err != nil {
		log.Fatalf("bidirectionalStreamRequest.err: %v", err)
	}

}

func responseStreamRequest(client hellowstream.StreamGreeterClient, request *hellowstream.StreamRequest) error {
	stream, err := client.ResponseStream(context.Background(), request)
	if err != nil {
		return err
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Infof("response: %s", recv.Message)
	}
	return nil
}

func clientStreamRequest(client hellowstream.StreamGreeterClient, request *hellowstream.StreamRequest) error {
	return nil
}

func bidirectionalStreamRequest(client hellowstream.StreamGreeterClient, request *hellowstream.StreamRequest) error {
	return nil
}