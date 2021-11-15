package main

import (
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "log"
    "net/http"
    _ "tomgs-go/learning-grpc-gateway/hello-world/api"
    "tomgs-go/learning-grpc-gateway/hello-world/route"
)

func main() {
    gwmux := runtime.NewServeMux()
    route.RegisterRoutes(gwmux)

    gwServer := &http.Server{
        Addr:    ":8090",
        Handler: gwmux,
    }

    log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
    log.Fatalln(gwServer.ListenAndServe())
}
