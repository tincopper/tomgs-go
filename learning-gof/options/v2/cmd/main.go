package main

import (
	"fmt"
	"time"
	v2 "tomgs-go/learning-gof/options/v2"
)

func main() {
	// Create a server with default options
	defaultServer := v2.NewServer()
	fmt.Printf("Default server: host=%s, port=%d, timeout=%v, ssl=%t, maxConn=%d\n",
		defaultServer.Host(), defaultServer.Port(), defaultServer.Timeout(),
		defaultServer.SSL(), defaultServer.MaxConn())

	// Create a server with some custom options
	customServer := v2.NewServer(
		v2.WithHost("example.com"),
		v2.WithPort(9000),
		v2.WithTimeout(60*time.Second),
		v2.WithSSL(),
	)
	fmt.Printf("Custom server: host=%s, port=%d, timeout=%v, ssl=%t, maxConn=%d\n",
		customServer.Host(), customServer.Port(), customServer.Timeout(),
		customServer.SSL(), customServer.MaxConn())

	// Create a server with all options customized
	fullCustomServer := v2.NewServer(
		v2.WithHost("api.mycompany.com"),
		v2.WithPort(443),
		v2.WithTimeout(120*time.Second),
		v2.WithSSL(),
		v2.WithMaxConn(1000),
	)
	fmt.Printf("Full custom server: host=%s, port=%d, timeout=%v, ssl=%t, maxConn=%d\n",
		fullCustomServer.Host(), fullCustomServer.Port(), fullCustomServer.Timeout(),
		fullCustomServer.SSL(), fullCustomServer.MaxConn())
}
