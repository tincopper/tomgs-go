package v1

import (
	"fmt"
	"reflect"
)

// GenericDynamicProxy 泛型动态代理结构体
type GenericDynamicProxy[T any] struct {
	target   *T               // 被代理对象
	handlers []HandlerFunc[T] // 处理器链
}

// HandlerFunc 代理处理器函数类型
type HandlerFunc[T any] func(method string, args []any, next func() []any) []any

// GenericDynamicProxy 创建动态代理实例
func NewGenericDynamicProxy[T any](target *T, handlers ...HandlerFunc[T]) *GenericDynamicProxy[T] {
	return &GenericDynamicProxy[T]{
		target:   target,
		handlers: handlers,
	}
}

// Invoke 方法调用入口
func (p *GenericDynamicProxy[T]) Invoke(method string, args ...any) []any {
	// 构建调用链
	var next func() []any
	next = func() []any {
		// 最终执行实际方法
		targetValue := reflect.ValueOf(p.target)
		methodValue := targetValue.MethodByName(method)

		if !methodValue.IsValid() {
			panic(fmt.Sprintf("method %s not found", method))
		}

		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}

		results := methodValue.Call(in)
		out := make([]any, len(results))
		for i, res := range results {
			out[i] = res.Interface()
		}
		return out
	}

	// 按顺序执行处理器链
	for i := len(p.handlers) - 1; i >= 0; i-- {
		next = func(currentNext func() []any, handler HandlerFunc[T]) func() []any {
			return func() []any {
				return handler(method, args, currentNext)
			}
		}(next, p.handlers[i])
	}

	return next()
}
