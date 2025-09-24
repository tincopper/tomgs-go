# Options Design Pattern in Go

This project demonstrates the Options design pattern in Go, which is a functional approach to configuring structs with many optional fields.

## What is the Options Pattern?

The Options pattern is a Go design pattern that provides a clean and flexible way to configure objects with many optional parameters. Instead of having multiple constructor functions or requiring a config struct, the Options pattern uses functional options to allow users to selectively configure only the parameters they care about.

## Implementation Details

The implementation consists of:

1. A `Server` struct that represents a server configuration
2. A `ServerOptions` struct that holds all possible configuration options
3. An `Option` type which is a function that modifies `ServerOptions`
4. A `NewServer` constructor that accepts variadic `Option` parameters
5. Several `WithX` functions that return `Option` functions for each configurable field

## Usage

```go
// Create a server with default options
defaultServer := options.NewServer()

// Create a server with some custom options
customServer := options.NewServer(
    options.WithHost("example.com"),
    options.WithPort(9000),
    options.WithTimeout(60*time.Second),
    options.WithSSL(),
)

// Create a server with all options customized
fullCustomServer := options.NewServer(
    options.WithHost("api.mycompany.com"),
    options.WithPort(443),
    options.WithTimeout(120*time.Second),
    options.WithSSL(),
    options.WithMaxConn(1000),
)
```

## Benefits

1. **Clean API**: Users can easily see what options are available
2. **Extensible**: New options can be added without breaking existing code
3. **Default Values**: Sensible defaults are provided automatically
4. **Type Safety**: All options are type-checked at compile time
5. **Readability**: The configuration is self-documenting

## Running the Example

To run the example:

```bash
go run cmd/main.go
```

## Running Tests

To run the unit tests:

```bash
go test -v
```

## Running Benchmarks

To run the benchmarks:

```bash
go test -bench=.
```