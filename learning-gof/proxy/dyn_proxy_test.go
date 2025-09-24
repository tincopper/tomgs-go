package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

type DynProxy struct {
	target interface{}
}

func (p *DynProxy) Invoke(methodName string, args ...interface{}) []reflect.Value {
	t := reflect.ValueOf(p.target)
	method := t.MethodByName(methodName)
	if !method.IsValid() {
		panic("Method not found")
	}

	// 前置处理
	fmt.Printf("DynamicProxy: calling %s\n", methodName)

	// 调用实际方法
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	result := method.Call(in)

	// 后置处理
	fmt.Println("DynamicProxy: call completed")

	return result
}

// 使用示例
type Service struct{}

func (s *Service) Process(data string) string {
	return "Processing " + data
}

func TestDynamicProxy(test *testing.T) {
	service := &Service{}
	proxy := &DynProxy{target: service}
	ret := proxy.Invoke("Process", "test data")
	fmt.Println(ret[0].String())
}
