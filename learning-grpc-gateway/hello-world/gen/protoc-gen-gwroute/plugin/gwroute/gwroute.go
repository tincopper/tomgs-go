package gwroute

import (
    "fmt"
)

import (
    "github.com/golang/protobuf/protoc-gen-go/generator"
)

// generatedCodeVersion indicates a version of the generated codes.
// It is incremented whenever an incompatibility between the generated codes and
// the grpc package is introduced; the generated codes references
// a constant, grpc.SupportPackageIsVersionN (where N is generatedCodeVersion).
const generatedCodeVersion = 4

// Paths for packages used by codes generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
    contextPkgPath = "context"
    grpcPkgPath    = "google.golang.org/grpc"
    codePkgPath    = "google.golang.org/grpc/codes"
    statusPkgPath  = "google.golang.org/grpc/status"
)

func init() {
    generator.RegisterPlugin(new(gwroute))
}

// triple is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for gRPC-triple support.
type gwroute struct {
    gen *generator.Generator
}

func (g *gwroute) Init(gen *generator.Generator) {
    fmt.Print("Init gwroute plugin...")
    g.gen = gen
}

func (g *gwroute) Generate(file *generator.FileDescriptor) {
    if len(file.FileDescriptorProto.Service) == 0 {
        return
    }
    // func init() {
    //    route.AddRoute(GreeterRouteDefinition)
    //}
    //
    //func GreeterRouteDefinition(gwmux *runtime.ServeMux) {
    //    // Create a client connection to the gRPC server we just started
    //    // This is where the gRPC-Gateway proxies the requests
    //    route.AddUnitRoute("0.0.0.0:8080", gwmux, func(gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {
    //        return RegisterGreeterHandler(context.Background(), gwmux, conn)
    //    })
    //}

    g.P("func init() {")
    for _, serviceDescriptorProto := range file.FileDescriptorProto.Service {
        serviceName := serviceDescriptorProto.GetName()
        g.P(fmt.Sprintf("route.AddRoute(%sRouteDefinition)", serviceName))
    }
    g.P("}")

    for _, serviceDescriptorProto := range file.FileDescriptorProto.Service {
        serviceName := serviceDescriptorProto.GetName()
        g.P(fmt.Sprintf("func %sRouteDefinition(gwmux *runtime.ServeMux) {", serviceName))
        g.P("\troute.AddUnitRoute(\"0.0.0.0:8080\", gwmux, func(gwmux *runtime.ServeMux, conn *grpc.ClientConn) error {")
        g.P(fmt.Sprintf("\t\treturn Register%sHandler(context.Background(), gwmux, conn)", serviceName))
        g.P("\t})")
        g.P("}")
    }

}

func (g *gwroute) GenerateImports(file *generator.FileDescriptor) {
    // import (
    //    "context"
    //    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    //    "google.golang.org/grpc"
    //    "tomgs-go/learning-grpc-gateway/hello-world/route"
    //)
    g.P("import (")
    g.P(`"context"`)
    g.P(`"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"`)
    g.P(`"google.golang.org/grpc"`)
    g.P(`"tomgs-go/learning-grpc-gateway/hello-world/route"`)
    g.P(` ) `)
}

// Name returns the name of this plugin, "grpc".
func (g *gwroute) Name() string {
    return "gwroute"
}

// P forwards to g.gen.P.
func (g *gwroute) P(args ...interface{}) {
    g.gen.P(args...)
}
