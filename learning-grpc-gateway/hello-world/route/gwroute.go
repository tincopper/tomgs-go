package route

import (
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "log"
)

var routes []func(gwmux *runtime.ServeMux)

func RegisterRoutes(gwmux *runtime.ServeMux) {
    //api.AddRoutes(gwmux)
    for _, routeFunc := range routes {
        if routeFunc != nil {
            routeFunc(gwmux)
        }
    }
}

func AddRoute(routeFunc func(gwmux *runtime.ServeMux)) {
    routes = append(routes, routeFunc)
}

func AddUnitRoute(target string, gwmux *runtime.ServeMux, route func(gwmux *runtime.ServeMux, conn *grpc.ClientConn) error) {
    // Create a client connection to the gRPC server we just started
    // This is where the gRPC-Gateway proxies the requests
    conn, err := grpc.DialContext(
        context.Background(),
        target, //pod service name
        grpc.WithBlock(),
        grpc.WithInsecure(),
    )
    if err != nil {
        log.Fatalln("Failed to dial server:", err)
    }
    // Register Greeter
    //err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
    err = route(gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }
}