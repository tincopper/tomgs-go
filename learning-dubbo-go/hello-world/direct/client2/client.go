/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

var tripleGreeterImpl = new(api.GreeterClientImpl)
var tripleGreeter2Impl = new(api.Greeter2ClientImpl)

// 错误演示
func init() {
	config.SetConsumerService(tripleGreeterImpl)
	config.SetConsumerService(tripleGreeter2Impl)
	refConf := config.ReferenceConfig{
		Protocol:      "tri",
		URL:           "tri://127.0.0.1:20000",
	}
	rootConfig := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			AddReference("GreeterClientImpl", &refConf).
			Build()).
		//AddRegistry("zkRegistryKey", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
		Build()

	if err := config.Load(config.WithRootConfig(rootConfig)); err != nil {
		panic(err)
	}

	///////////////////////////////////////////////

	// 这个不能用同一个对象
	refConf2 := config.ReferenceConfig{
		Protocol:      "tri",
		URL:           "tri://127.0.0.1:20000",
	}

	rootConfig2 := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			AddReference("Greeter2ClientImpl", &refConf2).
			Build()).
		//AddRegistry("zkRegistryKey", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
		Build()

	// Load只会执行一次
	if err := config.Load(config.WithRootConfig(rootConfig2)); err != nil {
		panic(err)
	}
}

// There is no need to export DUBBO_GO_CONFIG_PATH, as you are using config api to set config
func main() {
	logger.Info("start to test dubbo")
	helloGreeter()
	helloGreeter2()
}

func helloGreeter() {
	req := &api.HelloRequest{
		Name: "laurence",
	}
	reply, err := tripleGreeterImpl.SayHello(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("client response result: %v", reply)
}

func helloGreeter2() {
	req := &api.HelloRequest2{
		Name: "laurence2",
	}
	reply, err := tripleGreeter2Impl.SayHello2(context.Background(), req)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("client2 response result: %v\n", reply)
}