package protoreflect

import (
    "fmt"
    "github.com/golang/protobuf/proto"
    "github.com/jhump/protoreflect/desc/protoparse"
    "log"
    test "tomgs-go/learning-protoc/protoc-plugin/protoc-gen-gwroute2/example/service"
)

// 解析proto文件
func ParseProto() {
    parser := protoparse.Parser{
        ImportPaths: []string{"E:\\go_workspace\\src\\tomgs-go\\learning-protoc\\protoc-plugin\\protoc-gen-gwroute2\\example\\",
            "E:\\go_workspace\\src\\tomgs-go\\learning-protoc\\protoc-plugin\\protoc-gen-gwroute2\\example\\googleapis"},
    }
    files, err := parser.ParseFiles("helloworld.proto")
    if err != nil {
        log.Fatal("parse error ", err)
    }
    for _, fileDescriptor := range files {
        fmt.Println(fileDescriptor.GetName())
        str := fileDescriptor.GetOptions().String()
        fmt.Println(str)

        rawFields := proto.MessageV2(fileDescriptor.GetOptions()).ProtoReflect().GetUnknown()
        endpointValue := string(rawFields)
        fmt.Println(endpointValue)

        value, _ := proto.GetExtension(fileDescriptor.GetOptions(), test.E_Endpoint)
        fmt.Println(fmt.Sprint(*value.(*string)))

        endpointExtension := fileDescriptor.FindExtension("google.protobuf.FileOptions", 51234)
        fmt.Println(endpointExtension.GetName())

        for _, extension := range fileDescriptor.GetExtensions() {
            name := extension.GetName()
            fmt.Println(name)
            str := extension.GetOptions().String()
            fmt.Println(str)
        }
    }
}
