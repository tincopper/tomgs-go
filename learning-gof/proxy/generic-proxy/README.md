# Generic Dynamic Proxy in Go

This implementation provides a generic dynamic proxy mechanism in Go with support for pre/post processing logic through interceptors.

## Features

1. **Generic Proxy**: Works with any Go interface and implementation
2. **Interceptor Pattern**: Supports before/after processing logic
3. **Error Handling**: Properly handles and propagates errors
4. **Flexible Design**: Allows chaining of multiple interceptors
5. **Type Safety**: Maintains type safety through reflection

## Components

### Proxy
The main proxy struct that holds the target object and interceptor.

### Interceptor
Interface that defines before/after processing methods:
- `Before(mi *MethodInvocation) error`: Called before the method execution
- `After(mi *MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error)`: Called after the method execution

### MethodInvocation
Represents a method call with its arguments and provides a way to proceed with the actual method execution.

## Usage

1. Create your target object that implements an interface
2. Implement the `Interceptor` interface for your pre/post processing logic
3. Create a proxy using `genericproxy.NewProxy(target, interceptor)`
4. Get proxied methods using `proxy.GetMethod("MethodName")` and cast to the appropriate function type
5. Call the proxied methods as you would call the original methods

## Examples

See `example.go` for a basic example with logging interceptor and `advanced_example.go` for a more complex example with validation and timing interceptors.

## Running the Examples

```bash
# Basic example
go run example.go

# Advanced example
go run advanced_example.go
```

## Performance

The proxy implementation uses reflection, which adds some overhead compared to direct method calls:

```
BenchmarkDirectCall-16     1000000000    0.330 ns/op    0 B/op    0 allocs/op
BenchmarkProxiedCall-16     1871569      620.8 ns/op    200 B/op  4 allocs/op
```

The proxied call is approximately 1800x slower than a direct call, with 200 bytes of allocations.
This overhead is expected due to the use of reflection and should be considered when using this implementation.