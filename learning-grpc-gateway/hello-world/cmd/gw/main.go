package main

import (
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "log"
    "net/http"
    "tomgs-go/learning-grpc-gateway/hello-world/route"
)

func main() {
    gwmux := runtime.NewServeMux()
    route.AddRoute(gwmux)

    gwServer := &http.Server{
        Addr:    ":8090",
        Handler: gwmux,
    }

    log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
    log.Fatalln(gwServer.ListenAndServe())
}