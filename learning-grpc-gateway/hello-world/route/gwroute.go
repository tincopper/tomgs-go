package route

import (
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "log"
    helloworldpb "tomgs-go/learning-grpc-gateway/hello-world/api"
)

func AddRoute(gwmux *runtime.ServeMux) {
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
    err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }

    conn, err = grpc.DialContext(
        context.Background(),
        "0.0.0.0:8081",
        grpc.WithBlock(),
        grpc.WithInsecure(),
    )
    if err != nil {
        log.Fatalln("Failed to dial server:", err)
    }
    // Register Greeter
    err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }

}