package v2

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

// 定义业务服务
type UserService struct{}

func (s *UserService) GetUser(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("invalid user id")
	}
	return fmt.Sprintf("User-%d", id), nil
}

func (s *UserService) CreateUser(name string) (int, error) {
	if len(name) < 3 {
		return 0, fmt.Errorf("name too short")
	}
	return len(name), nil
}

func TestGenericProxy(t *testing.T) {
	// 创建被代理对象
	userService := &UserService{}

	// 创建动态代理并添加中间件
	proxy := NewDynamicProxy(
		userService,
		LoggingMiddleware[*UserService],
		HandlerFunc[UserService](TimeoutMiddleware[*UserService](time.Second)),
		HandlerFunc[UserService](CircuitBreakerMiddleware[*UserService](3, time.Minute)),
	)
	defer proxy.Close()

	// 方式1：直接调用
	results, err := proxy.Invoke(context.Background(), "GetUser", 123)
	if err != nil {
		log.Fatal(err)
	}
	user := results[0].(string)
	fmt.Println("GetUser result:", user)

	// 方式2：生成类型安全包装器
	wrappedService, err := proxy.GenerateWrapper()
	if err != nil {
		log.Fatal(err)
	}

	// 像调用原始服务一样使用
	id, err := wrappedService.CreateUser("Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateUser result:", id)
}
