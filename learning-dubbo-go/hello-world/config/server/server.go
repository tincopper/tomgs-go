package main

import (
    "context"
)

import (
    "dubbo.apache.org/dubbo-go/v3/common/logger"
    "dubbo.apache.org/dubbo-go/v3/config"
    _ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
    "tomgs-go/learning-dubbo-go/hello-world/api"
)

func main() {
    config.SetProviderService(&GreeterProvider{})

    rc := config.NewRootConfigBuilder().
        SetProvider(config.NewProviderConfigBuilder().
            AddService("GreeterProvider", config.NewServiceConfigBuilder().
                SetInterface("org.apache.dubbo.UserProvider").
                SetProtocolIDs("tripleProtocolKey").
                Build()).
            SetRegistryIDs("registryKey").
            Build()).
        AddProtocol("tripleProtocolKey", config.NewProtocolConfigBuilder().
            SetName("tri").
            Build()).
        AddRegistry("registryKey", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
        Build()

    if err := rc.Init(); err != nil {
        panic(err)
    }

    select {}
}

type GreeterProvider struct {
    api.GreeterServer
}

func (s *GreeterProvider) SayHelloStream(svr api.Greeter_SayHelloStreamServer) error {
    c, err := svr.Recv()
    if err != nil {
        return err
    }
    logger.Infof("Dubbo-go3 GreeterProvider recv 1 user, name = %s\n", c.Name)
    c2, err := svr.Recv()
    if err != nil {
        return err
    }
    logger.Infof("Dubbo-go3 GreeterProvider recv 2 user, name = %s\n", c2.Name)
    c3, err := svr.Recv()
    if err != nil {
        return err
    }
    logger.Infof("Dubbo-go3 GreeterProvider recv 3 user, name = %s\n", c3.Name)

    svr.Send(&api.User{
        Name: "hello " + c.Name,
        Age:  18,
        Id:   "123456789",
    })
    svr.Send(&api.User{
        Name: "hello " + c2.Name,
        Age:  19,
        Id:   "123456789",
    })
    return nil
}

func (s *GreeterProvider) SayHello(ctx context.Context, in *api.HelloRequest) (*api.User, error) {
    logger.Infof("Dubbo3 GreeterProvider get user name = %s\n", in.Name)
    return &api.User{Name: "Hello " + in.Name, Id: "12345", Age: 21}, nil
}