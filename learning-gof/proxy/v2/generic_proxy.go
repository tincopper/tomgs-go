package v2

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

// DynamicProxy 泛型动态代理核心结构
type DynamicProxy[T any] struct {
	target      *T                 // 被代理对象
	handlers    []HandlerFunc[T]   // 处理器链
	methodCache sync.Map           // 方法缓存
	ctx         context.Context    // 上下文
	cancel      context.CancelFunc // 取消函数
}

// HandlerFunc 代理处理器函数类型
type HandlerFunc[T any] func(
	ctx context.Context,
	method string,
	args []any,
	next func(context.Context) ([]any, error),
) ([]any, error)

// NewDynamicProxy 创建动态代理实例
func NewDynamicProxy[T any](
	target *T,
	handlers ...HandlerFunc[T],
) *DynamicProxy[T] {
	ctx, cancel := context.WithCancel(context.Background())
	return &DynamicProxy[T]{
		target:   target,
		handlers: handlers,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Close 释放代理资源
func (p *DynamicProxy[T]) Close() {
	p.cancel()
}

// Invoke 安全调用方法
func (p *DynamicProxy[T]) Invoke(
	ctx context.Context,
	method string,
	args ...any,
) (results []any, err error) {
	// 合并上下文
	ctx, cancel := context.WithCancel(p.ctx)
	defer cancel()

	// 异常恢复
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("proxy panic: %v", r)
		}
	}()

	// 构建调用链
	var next func(context.Context) ([]any, error)
	next = func(ctx context.Context) ([]any, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// 获取缓存的方法
			methodValue, err := p.getMethod(method)
			if err != nil {
				return nil, err
			}

			// 转换参数
			in, err := p.convertArgs(methodValue, args)
			if err != nil {
				return nil, err
			}

			// 执行实际方法
			results := methodValue.Call(in)
			return p.convertResults(results), nil
		}
	}

	// 反向组装中间件链
	for i := len(p.handlers) - 1; i >= 0; i-- {
		next = func(
			currentNext func(context.Context) ([]any, error),
			handler HandlerFunc[T],
		) func(context.Context) ([]any, error) {
			return func(ctx context.Context) ([]any, error) {
				return handler(ctx, method, args, currentNext)
			}
		}(next, p.handlers[i])
	}

	return next(ctx)
}

// GenerateWrapper 生成类型安全的包装器
func (p *DynamicProxy[T]) GenerateWrapper() (*T, error) {
	var target T
	t := reflect.TypeOf(&target).Elem()

	value := reflect.New(t)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		wrapper := p.createMethodWrapper(method.Name)
		value.Elem().FieldByName(method.Name).Set(
			reflect.MakeFunc(method.Type, wrapper),
		)
	}

	return value.Interface().(*T), nil
}

// 内部工具方法
func (p *DynamicProxy[T]) getMethod(method string) (reflect.Value, error) {
	// 检查缓存
	if cached, ok := p.methodCache.Load(method); ok {
		return cached.(reflect.Value), nil
	}

	targetValue := reflect.ValueOf(p.target)
	methodValue := targetValue.MethodByName(method)
	if !methodValue.IsValid() {
		return reflect.Value{}, fmt.Errorf("method %s not found", method)
	}

	p.methodCache.Store(method, methodValue)
	return methodValue, nil
}

func (p *DynamicProxy[T]) convertArgs(
	method reflect.Value,
	args []any,
) ([]reflect.Value, error) {
	methodType := method.Type()
	if len(args) != methodType.NumIn() {
		return nil, fmt.Errorf(
			"argument count mismatch: expected %d, got %d",
			methodType.NumIn(),
			len(args),
		)
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValue := reflect.ValueOf(arg)
		paramType := methodType.In(i)

		if !argValue.Type().AssignableTo(paramType) {
			return nil, fmt.Errorf(
				"argument %d type mismatch: expected %v, got %v",
				i,
				paramType,
				argValue.Type(),
			)
		}

		in[i] = argValue
	}

	return in, nil
}

func (p *DynamicProxy[T]) convertResults(results []reflect.Value) []any {
	out := make([]any, len(results))
	for i, res := range results {
		out[i] = res.Interface()
	}
	return out
}

func (p *DynamicProxy[T]) createMethodWrapper(methodName string) func([]reflect.Value) []reflect.Value {
	return func(args []reflect.Value) []reflect.Value {
		// 转换输入参数
		in := make([]any, len(args))
		for i, arg := range args {
			in[i] = arg.Interface()
		}

		// 调用代理
		results, err := p.Invoke(context.Background(), methodName, in...)
		if err != nil {
			// 处理错误返回
			errVal := reflect.ValueOf(err)
			errType := reflect.TypeOf((*error)(nil)).Elem()

			// 检查方法是否有error返回值
			method, _ := p.getMethod(methodName)
			if method.Type().NumOut() > 0 &&
				method.Type().Out(method.Type().NumOut()-1).Implements(errType) {
				out := make([]reflect.Value, method.Type().NumOut())
				for i := 0; i < len(out)-1; i++ {
					out[i] = reflect.Zero(method.Type().Out(i))
				}
				out[len(out)-1] = errVal
				return out
			}

			panic(err)
		}

		// 转换返回结果
		out := make([]reflect.Value, len(results))
		for i, res := range results {
			out[i] = reflect.ValueOf(res)
		}
		return out
	}
}

/*************************** 预定义中间件 ***************************/

// LoggingMiddleware 日志记录中间件
func LoggingMiddleware[T any](
	ctx context.Context,
	method string,
	args []any,
	next func(context.Context) ([]any, error),
) ([]any, error) {
	start := time.Now()
	log.Printf("Calling %s with args: %v", method, args)

	results, err := next(ctx)

	duration := time.Since(start)
	if err != nil {
		log.Printf("Method %s failed after %v: %v", method, duration, err)
	} else {
		log.Printf("Method %s executed in %v, result: %v", method, duration, results)
	}

	return results, err
}

// TimeoutMiddleware 超时控制中间件
func TimeoutMiddleware[T any](timeout time.Duration) HandlerFunc[T] {
	return func(
		ctx context.Context,
		method string,
		args []any,
		next func(context.Context) ([]any, error),
	) ([]any, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		done := make(chan struct{})
		var results []any
		var err error

		go func() {
			defer close(done)
			results, err = next(ctx)
		}()

		select {
		case <-done:
			return results, err
		case <-ctx.Done():
			return nil, fmt.Errorf("method %s timed out after %v", method, timeout)
		}
	}
}

// CircuitBreakerMiddleware 熔断器中间件
func CircuitBreakerMiddleware[T any](
	maxFailures int,
	resetTimeout time.Duration,
) HandlerFunc[T] {
	var (
		failureCount int
		lastFailure  time.Time
		mu           sync.Mutex
	)

	return func(
		ctx context.Context,
		method string,
		args []any,
		next func(context.Context) ([]any, error),
	) ([]any, error) {
		mu.Lock()
		defer mu.Unlock()

		// 检查是否需要重置计数器
		if time.Since(lastFailure) > resetTimeout {
			failureCount = 0
		}

		// 检查是否触发熔断
		if failureCount >= maxFailures {
			return nil, fmt.Errorf("service unavailable (circuit breaker tripped)")
		}

		results, err := next(ctx)
		if err != nil {
			failureCount++
			lastFailure = time.Now()
		}

		return results, err
	}
}
