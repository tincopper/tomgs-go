package api

import (
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "tomgs-go/learning-grpc-gateway/hello-world/route"
)

func init() {
    route.AddRoute(GreeterRouteDefinition)
}

func GreeterRouteDefinition(gwmux *runtime.ServeMux) {
    // Create a client connection to the gRPC server we just started
    // This is where the gRPC-Gateway proxies the requests
    route.AddUnitRoute("0.0.0.0:8080", gwmux, func(gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
        return RegisterGreeterHandler(context.Background(), gwmux, conn)
    })
}
