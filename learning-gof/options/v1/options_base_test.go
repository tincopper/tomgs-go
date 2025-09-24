package v1

import (
	"testing"
	"time"
)

func TestOptionsBase(t *testing.T) {
	// 使用默认配置
	svc1 := NewService()
	svc1.Run()

	// 使用自定义配置
	svc2 := NewService(
		WithTimeout(30*time.Second),
		WithMaxRetries(5),
		WithCache(true),
	)
	svc2.Run()

	// 使用自定义Logger
	customLogger := &defaultLogger{}
	svc3 := NewService(
		WithLogger(customLogger),
		WithMaxRetries(1),
	)
	svc3.Run()
}
