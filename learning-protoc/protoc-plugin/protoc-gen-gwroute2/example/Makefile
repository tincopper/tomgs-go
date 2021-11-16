# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 安装插件
gwroute-install:
	go install ../../protoc-gen-gwroute2

endpoint-gen:
	protoc -I ./ -I ./googleapis/ endpoint.proto  --go_out ./service --go_opt=paths=source_relative

go-gen:
	protoc -I ./ -I ./googleapis/ helloworld.proto  --go_out .

grpc-gen:
	protoc -I ./ -I ./googleapis/ helloworld.proto  --go-grpc_out .

gateway-gen:
	protoc -I ./ -I ./googleapis/ helloworld.proto --grpc-gateway_out .

gwroute-gen:
	protoc -I ./ -I ./googleapis/ helloworld.proto  --gwroute2_out=plugins=gwroute2:.
