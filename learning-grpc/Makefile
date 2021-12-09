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

plugin-install:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

build-hello-world-old:
	if not exist genproto\customapis\endpoint mkdir genproto\customapis\endpoint
	protoc -I ./hello-world/api -I ./third_party/googleapis/ -I ./third_party/customapis/ endpoint.proto --go_out genproto/customapis/endpoint --go_opt=paths=source_relative
	protoc -I ./hello-world/api -I ./third_party/googleapis/ -I ./third_party/customapis/ helloworld.proto  --go_out ./hello-world/api --go_opt=paths=source_relative
	protoc -I ./hello-world/api -I ./third_party/googleapis/ -I ./third_party/customapis/ helloworld.proto  --go-grpc_out ./hello-world/api


build-hello-world:
	protoc -I ./proto/ -I ./third_party/googleapis/ ./proto/helloworld.proto  --go_out ./hello-world/api
	protoc -I ./proto/ -I ./third_party/googleapis/ ./proto/helloworld.proto  --go-grpc_out ./hello-world/api

build-hello-stream:
	protoc -I ./proto/ -I ./third_party/googleapis/ ./proto/hellostream.proto  --go_out ./hello-stream/api
	protoc -I ./proto/ -I ./third_party/googleapis/ ./proto/hellostream.proto  --go-grpc_out ./hello-stream/api

build: build-hello-world build-hello-stream
