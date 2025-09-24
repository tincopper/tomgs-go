package v1

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// 定义业务服务
type UserService struct{}

func (s *UserService) GetUser(id int) (string, error) {
	return fmt.Sprintf("User-%d", id), nil
}

func (s *UserService) CreateUser(name string) int {
	return len(name)
}

// 日志处理器
func loggingMiddleware[T any](method string, args []any, next func() []any) []any {
	start := time.Now()
	log.Printf("Calling %s with args: %v", method, args)

	result := next()

	log.Printf("Method %s executed in %v, result: %v",
		method, time.Since(start), result)
	return result
}

// 缓存处理器
func cacheMiddleware[T any](method string, args []any, next func() []any) []any {
	cacheKey := fmt.Sprintf("%s%v", method, args)

	// 这里简化实现，实际应该用真正的缓存
	if cached, ok := cache[cacheKey]; ok {
		return cached
	}

	result := next()
	cache[cacheKey] = result
	return result
}

var cache = make(map[string][]any)

func TestDynGenericProxy(t *testing.T) {
	// 创建被代理对象
	userService := &UserService{}

	// 创建动态代理
	proxy := NewGenericDynamicProxy(
		userService,
		loggingMiddleware[*UserService],
		cacheMiddleware[*UserService],
	)

	// 调用方法
	ret := proxy.Invoke("GetUser", 123)
	fmt.Println("GetUser result:", ret[0], ret[1])

	ret = proxy.Invoke("CreateUser", "Alice")
	fmt.Println("CreateUser result:", ret[0])
}
