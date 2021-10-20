package main

import (
    "context"
    "dubbo.apache.org/dubbo-go/v3/common/constant"
    "dubbo.apache.org/dubbo-go/v3/config"
    "dubbo.apache.org/dubbo-go/v3/config/generic"
    "dubbo.apache.org/dubbo-go/v3/protocol/dubbo"
    hessian "github.com/apache/dubbo-go-hessian2"
    gxlog "github.com/dubbogo/gost/log"
    "time"
)

var (
    appName        = "UserConsumer"
    registryConfig = &config.RegistryConfig{
        Protocol: "zookeeper",
        Address:  "127.0.0.1:2181",
    }
    referenceConfig = config.ReferenceConfig{
        InterfaceName: "com.tomgs.learning.dubbo.api.IGreeter",
        Cluster:       "failover",
        // registry需要配置文件
        RegistryIDs: []string{"zk"},
        Protocol:    dubbo.DUBBO,
        Generic:     "true",
    }
)

func init() {
    rootConfig := config.NewRootConfigBuilder().
        AddRegistry("zk", registryConfig).
        Build()
    _ = rootConfig.Init()
    _ = referenceConfig.Init(rootConfig)
    referenceConfig.GenericLoad(appName) //appName is the unique identification of RPCService
    time.Sleep(1 * time.Second)
}

// need to setup environment variable "CONF_CONSUMER_FILE_PATH" to "conf/client.yml" before run
func main() {
    call()
}

func call() {
    // 设置attachment
    ctx := context.WithValue(context.TODO(), constant.AttachmentKey, map[string]string{"tag": "test"})

    resp, err := referenceConfig.GetRPCService().(*generic.GenericService).Invoke(
        ctx,
        "sayHello",
        []string{"com.tomgs.learning.dubbo.api.Helloworld$HelloRequest"},
        []hessian.Object{
            map[string]string{"name": "roshi"},
        },
    )
    if err != nil {
        panic(err)
    }
    gxlog.CInfo("success called res: %+v\n", resp)
}
