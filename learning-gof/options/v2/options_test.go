package v2

import (
	"testing"
	"time"
)

func TestNewServerWithDefaults(t *testing.T) {
	server := NewServer()

	if server.Host() != "localhost" {
		t.Errorf("Expected host to be 'localhost', got '%s'", server.Host())
	}

	if server.Port() != 8080 {
		t.Errorf("Expected port to be 8080, got %d", server.Port())
	}

	if server.Timeout() != 30*time.Second {
		t.Errorf("Expected timeout to be 30s, got %v", server.Timeout())
	}

	if server.SSL() != false {
		t.Errorf("Expected SSL to be false, got %t", server.SSL())
	}

	if server.MaxConn() != 100 {
		t.Errorf("Expected MaxConn to be 100, got %d", server.MaxConn())
	}
}

func TestNewServerWithOptions(t *testing.T) {
	server := NewServer(
		WithHost("example.com"),
		WithPort(9000),
		WithTimeout(60*time.Second),
		WithSSL(),
		WithMaxConn(500),
	)

	if server.Host() != "example.com" {
		t.Errorf("Expected host to be 'example.com', got '%s'", server.Host())
	}

	if server.Port() != 9000 {
		t.Errorf("Expected port to be 9000, got %d", server.Port())
	}

	if server.Timeout() != 60*time.Second {
		t.Errorf("Expected timeout to be 60s, got %v", server.Timeout())
	}

	if server.SSL() != true {
		t.Errorf("Expected SSL to be true, got %t", server.SSL())
	}

	if server.MaxConn() != 500 {
		t.Errorf("Expected MaxConn to be 500, got %d", server.MaxConn())
	}
}

func TestNewServerWithPartialOptions(t *testing.T) {
	server := NewServer(
		WithHost("api.example.com"),
		WithPort(443),
	)

	if server.Host() != "api.example.com" {
		t.Errorf("Expected host to be 'api.example.com', got '%s'", server.Host())
	}

	if server.Port() != 443 {
		t.Errorf("Expected port to be 443, got %d", server.Port())
	}

	// Other options should still have default values
	if server.Timeout() != 30*time.Second {
		t.Errorf("Expected timeout to be 30s, got %v", server.Timeout())
	}

	if server.SSL() != false {
		t.Errorf("Expected SSL to be false, got %t", server.SSL())
	}

	if server.MaxConn() != 100 {
		t.Errorf("Expected MaxConn to be 100, got %d", server.MaxConn())
	}
}
