# 主要用于测试protoc插件的使用
# https://github.com/go-kiss/protoc-gen-markdown

plugin-install:
	go install \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/go-kiss/protoc-gen-markdown

# 使用的 protoc --version 为 libprotoc 3.18.1
# 使用的 protoc-gen-go --version 为 protoc-gen-go v1.27.1
# 使用的 protoc-gen-go-grpc --version 为 protoc-gen-go-grpc 1.1.0
# 使用的 protoc-gen-doc --version 为 protoc-gen-doc version 1.5.0
# 在根目录下执行 protoc 命令
gen-html:
	protoc -I ./third_party/googleapis/ -I ./proto/ \
	--go_out=api \
	--go-grpc_out=api \
	--doc_out=api \
	--doc_opt=html,index.html \
	proto/helloworld.proto

gen-md:
	protoc -I ./third_party/googleapis/ -I ./proto/ \
	--go_out=api \
	--go-grpc_out=api \
	--doc_out=api \
	--doc_opt=markdown,helloworld.md \
	proto/helloworld.proto

# set path prefix to /api
# protoc -I ./third_party/googleapis/ -I ./proto/ --markdown_out=Mhello.proto=./,prefix=/api:. ./proto/helloworld.proto
gen-md-by-go-kiss:
	protoc -I ./third_party/googleapis/ -I ./proto/ --markdown_out=Mhello.proto=./:./api ./proto/helloworld.proto