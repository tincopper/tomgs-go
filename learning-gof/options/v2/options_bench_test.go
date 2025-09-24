package v2

import (
	"testing"
	"time"
)

func BenchmarkNewServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewServer()
	}
}

func BenchmarkNewServerWithOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewServer(
			WithHost("example.com"),
			WithPort(9000),
			WithTimeout(60*time.Second),
			WithSSL(),
			WithMaxConn(500),
		)
	}
}
