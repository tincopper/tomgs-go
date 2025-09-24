package genericproxy

import (
	"testing"
)

func BenchmarkDirectCall(b *testing.B) {
	service := &mockService{value: 42}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.GetValue()
	}
}

func BenchmarkProxiedCall(b *testing.B) {
	service := &mockService{value: 42}
	interceptor := &SimpleInterceptor{}
	proxy := NewProxy(service, interceptor)
	getValueMethod := proxy.GetMethod("GetValue").(func() int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getValueMethod()
	}
}
