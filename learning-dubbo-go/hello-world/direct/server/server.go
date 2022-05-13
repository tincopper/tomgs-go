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

type GreeterProvider struct {
	api.UnimplementedGreeterServer
	//api.UnimplementedGreeter2Server
}

type Greeter2Provider struct {
	api.UnimplementedGreeter2Server
}

func (s *GreeterProvider) SayHello(ctx context.Context, in *api.HelloRequest) (*api.User, error) {
	logger.Infof("Dubbo3 GreeterProvider get user name = %s\n", in.Name)
	return &api.User{Name: "Hello " + in.Name, Id: "12345", Age: 21}, nil
}

func (s *Greeter2Provider) SayHello2(ctx context.Context, in *api.HelloRequest2) (*api.User2, error) {
	logger.Infof("Dubbo3 Greeter2Provider get user2 name = %s\n", in.Name)
	return &api.User2{Name: "Hello " + in.Name, Id: "123456", Age: 22}, nil
}

// There is no need to export DUBBO_GO_CONFIG_PATH, as you are using config api to set config
func main() {
	config.SetProviderService(&GreeterProvider{})
	config.SetProviderService(&Greeter2Provider{})

	rootConfig := config.NewRootConfigBuilder().
		SetProvider(config.NewProviderConfigBuilder().
			AddService("GreeterProvider", config.NewServiceConfigBuilder().
				Build()).
			AddService("Greeter2Provider", config.NewServiceConfigBuilder().
				Build()).
			Build()).
		//AddRegistry("zk", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
		AddProtocol("tripleKey", config.NewProtocolConfigBuilder().
			SetName("tri").
			SetIp("127.0.0.1").
			SetPort("20000").
			Build()).
		Build()

	if err := config.Load(config.WithRootConfig(rootConfig)); err != nil {
		panic(err)
	}

	select {}
}