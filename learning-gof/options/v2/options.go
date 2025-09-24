package v2

import (
	"time"
)

// Server represents a server configuration that we'll configure using the options pattern
type Server struct {
	host    string
	port    int
	timeout time.Duration
	ssl     bool
	maxConn int
}

// ServerOptions contains all the configuration options for a Server
type ServerOptions struct {
	Host    string
	Port    int
	Timeout time.Duration
	SSL     bool
	MaxConn int
}

// Option is a function that configures a ServerOptions
type Option func(*ServerOptions)

// NewServer creates a new Server with the provided options
func NewServer(opts ...Option) *Server {
	// Default options
	options := ServerOptions{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30 * time.Second,
		SSL:     false,
		MaxConn: 100,
	}

	// Apply the options
	for _, opt := range opts {
		opt(&options)
	}

	// Create and return the server
	return &Server{
		host:    options.Host,
		port:    options.Port,
		timeout: options.Timeout,
		ssl:     options.SSL,
		maxConn: options.MaxConn,
	}
}

// WithHost sets the host for the server
func WithHost(host string) Option {
	return func(o *ServerOptions) {
		o.Host = host
	}
}

// WithPort sets the port for the server
func WithPort(port int) Option {
	return func(o *ServerOptions) {
		o.Port = port
	}
}

// WithTimeout sets the timeout for the server
func WithTimeout(timeout time.Duration) Option {
	return func(o *ServerOptions) {
		o.Timeout = timeout
	}
}

// WithSSL enables SSL for the server
func WithSSL() Option {
	return func(o *ServerOptions) {
		o.SSL = true
	}
}

// WithMaxConn sets the maximum number of connections for the server
func WithMaxConn(maxConn int) Option {
	return func(o *ServerOptions) {
		o.MaxConn = maxConn
	}
}

// Getters for Server fields
func (s *Server) Host() string {
	return s.host
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) Timeout() time.Duration {
	return s.timeout
}

func (s *Server) SSL() bool {
	return s.ssl
}

func (s *Server) MaxConn() int {
	return s.maxConn
}
