package v1

import (
	"fmt"
	"log"
	"time"
)

// Logger 接口
type Logger interface {
	Log(msg string)
}

// 默认Logger实现
type defaultLogger struct{}

func (l *defaultLogger) Log(msg string) {
	log.Println(msg)
}

// Config 配置结构
type Config struct {
	Timeout     time.Duration
	MaxRetries  int
	Logger      Logger
	EnableCache bool
}

// Option 选项函数类型
type Option func(*Config)

// 各种选项函数
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithMaxRetries(retries int) Option {
	return func(c *Config) {
		if retries < 0 {
			panic("retries cannot be negative")
		}
		c.MaxRetries = retries
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

func WithCache(enable bool) Option {
	return func(c *Config) {
		c.EnableCache = enable
	}
}

// Service 服务结构
type Service struct {
	config *Config
}

// NewService 构造函数
func NewService(opts ...Option) *Service {
	// 默认配置
	config := &Config{
		Timeout:     10 * time.Second,
		MaxRetries:  3,
		Logger:      &defaultLogger{},
		EnableCache: false,
	}

	// 应用选项
	for _, opt := range opts {
		opt(config)
	}

	return &Service{config: config}
}

// Run 服务方法
func (s *Service) Run() {
	s.config.Logger.Log(fmt.Sprintf(
		"Service running with config: Timeout=%v, MaxRetries=%d, EnableCache=%v",
		s.config.Timeout,
		s.config.MaxRetries,
		s.config.EnableCache,
	))
}

// builder风格
type ServiceBuilder struct {
	config *Config
}

func NewBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		config: &Config{
			Timeout:    10 * time.Second,
			MaxRetries: 3,
			// 其他默认值...
		},
	}
}

func (b *ServiceBuilder) WithTimeout(timeout time.Duration) *ServiceBuilder {
	b.config.Timeout = timeout
	return b
}

func (b *ServiceBuilder) Build() *Service {
	return &Service{config: b.config}
}
