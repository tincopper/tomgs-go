package main

import (
    "context"
    "dubbo.apache.org/dubbo-go/v3/common/logger"
    "dubbo.apache.org/dubbo-go/v3/config"
    "dubbo.apache.org/dubbo-go/v3/config/generic"
    _ "dubbo.apache.org/dubbo-go/v3/imports"
    "dubbo.apache.org/dubbo-go/v3/protocol/dubbo"
    hessian "github.com/apache/dubbo-go-hessian2"
    "log"
    "net"
    "net/http"

    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"

    api "tomgs-go/learning-grpc-gateway/hello-world/api"
)

var dubboRefConf config.ReferenceConfig

func init() {
    dubboRefConf = newRefConf("kd.bos.debug.mservice.api.IGreeter", dubbo.DUBBO)
    //dubboRefConf = newRefConf("com.tomgs.learning.dubbo.api.IGreeter", dubbo.DUBBO)
    //dubboRefConf = newRefConf("kd.bos.service.DispatchService", dubbo.DUBBO)
}

type server struct {
    api.UnimplementedGreeterServer
}

//var greeterProvider = &api.GreeterClientImpl{}

//func init() {
//	config.SetConsumerService(greeterProvider)
//}

func NewServer() *server {
    return &server{}
}

func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.User, error) {
    name := callGetOneUser1(dubboRefConf, in.GetName())
    //name := callGetOneUser2(dubboRefConf, in.GetName())
    return &api.User{Name: name + " world"}, nil
}

func callGetOneUser1(refConf config.ReferenceConfig, arg string) string {
    resp, err := refConf.GetRPCService().(*generic.GenericService).Invoke(
        context.TODO(),
        "sayHello2",
        []string{"java.lang.String"},
        []hessian.Object{ arg },
    )
    if err != nil {
        panic(err)
    }
    logger.Infof("GetUser1(userId string) res: %+v", resp)
    return resp.(string)
}

func callGetOneUser2(refConf config.ReferenceConfig, arg string) string {
    resp, err := refConf.GetRPCService().(*generic.GenericService).Invoke(
        context.TODO(),
        "invoke",
        []string{"java.lang.String", "java.lang.String", "java.lang.String", "java.lang.Object[]"},
        []hessian.Object{
            "com.jdy.bd.assistant.servicehelper.ServiceFactory",
            "BD_BaseDataService",
            "getBaseDataList",
            "",
        },
    )
    if err != nil {
        //panic(err)
        logger.Infof(err.Error())
        return ""
    }
    logger.Infof("GetUser1(userId string) res: %+v", resp)
    return resp.(string)
}

func main() {
    // Create a listener on TCP port
    lis, err := net.Listen("tcp", ":8081")
    if err != nil {
        log.Fatalln("Failed to listen:", err)
    }

    // Create a gRPC server object
    s := grpc.NewServer()
    // Attach the Greeter service to the server
    api.RegisterGreeterServer(s, &server{})
    // Serve gRPC server
    log.Println("Serving gRPC on 0.0.0.0:8081")
    go func() {
        log.Fatalln(s.Serve(lis))
    }()

    // Create a client connection to the gRPC server we just started
    // This is where the gRPC-Gateway proxies the requests
    conn, err := grpc.DialContext(
        context.Background(),
        "0.0.0.0:8081",
        grpc.WithBlock(),
        grpc.WithInsecure(),
    )
    if err != nil {
        log.Fatalln("Failed to dial server:", err)
    }

    gwmux := runtime.NewServeMux()
    // Register Greeter
    err = api.RegisterGreeterHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }

    gwServer := &http.Server {
        Addr:    ":8090",
        Handler: gwmux,
    }

    log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
    log.Fatalln(gwServer.ListenAndServe())
}

const appName = "mservice"

func newRefConf(iface, protocol string) config.ReferenceConfig {
    /*registryConfig := &config.RegistryConfig{
        Protocol: "zookeeper",
        Address:  "127.0.0.1:2181",
    }*/

    refConf := config.ReferenceConfig{
        InterfaceName: iface,
        Cluster:       "failover",
        //RegistryIDs:   []string{"zk"},
        Protocol:      protocol,
        Generic:       "true",
        //Group:         "bd",
        //Version:       "1.0",
        //URL:           "dubbo://172.20.176.190:20880",
        URL:           "dubbo://127.0.0.1:50051",
    }

    rootConfig := config.NewRootConfigBuilder().
        //AddRegistry("zk", registryConfig).
        Build()
    _ = rootConfig.Init()
    _ = refConf.Init(rootConfig)
    refConf.GenericLoad(appName)

    return refConf
}
