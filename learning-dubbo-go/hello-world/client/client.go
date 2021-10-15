package main

import (
    "context"
    "time"
)

import (
    "dubbo.apache.org/dubbo-go/v3/common/logger"
    "dubbo.apache.org/dubbo-go/v3/config"
    _ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
    "tomgs-go/learning-dubbo-go/hello-world/api"
)

var greeterProvider = &api.GreeterClientImpl{}

func init() {
    config.SetConsumerService(greeterProvider)
}

func main() {
    // init rootConfig with config api
    rc := config.NewRootConfigBuilder().
        SetConsumer(config.NewConsumerConfigBuilder().
            SetRegistryIDs("zookeeper").
            AddReference("GreeterClientImpl", config.NewReferenceConfigBuilder().
                //SetInterface("org.apache.dubbo.UserProvider").
                SetInterface("com.tomgs.learning.dubbo.api.IGreeter").
                SetProtocol("dubbo").
                Build()).
            Build()).
        AddRegistry("zookeeper", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
        Build()

    // validate consumer greeterProvider
    if err := rc.Init(); err != nil {
        panic(err)
    }

    // waiting for service discovery
    time.Sleep(time.Second * 3)

    // run rpc invocation
    testSayHello()
}

func testSayHello() {
    ctx := context.Background()

    req := api.HelloRequest{
        Name: "tomgs",
    }
    user, err := greeterProvider.SayHello(ctx, &req)
    if err != nil {
        panic(err)
    }

    logger.Infof("Receive user = %+v\n", user)
}
