package main

import (
	"fmt"
	"log"
	"reflect"
	"tomgs-go/learning-gof/proxy/generic-proxy/genericproxy"
)

// Calculator is a simple interface for math operations
type Calculator interface {
	Add(a, b int) int
	Divide(a, b int) (int, error)
}

// calculator is the implementation of Calculator
type calculator struct{}

func (c *calculator) Add(a, b int) int {
	fmt.Printf("Calculating %d + %d\n", a, b)
	return a + b
}

func (c *calculator) Divide(a, b int) (int, error) {
	fmt.Printf("Calculating %d / %d\n", a, b)
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// LoggingInterceptor implements genericproxy.Interceptor for logging
type LoggingInterceptor struct{}

func (li *LoggingInterceptor) Before(mi *genericproxy.MethodInvocation) error {
	fmt.Printf("Before calling %s with args: %v\n", mi.MethodName, mi.Args)
	return nil
}

func (li *LoggingInterceptor) After(mi *genericproxy.MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	fmt.Printf("After calling %s, result: %v, error: %v\n", mi.MethodName, result, err)
	return result, err
}

func main() {
	// Create the target object
	calc := &calculator{}

	// Create the interceptor
	interceptor := &LoggingInterceptor{}

	// Create the proxy
	proxy := genericproxy.NewProxy(calc, interceptor)

	// Get proxied methods
	addMethod := proxy.GetMethod("Add").(func(int, int) int)
	divideMethod := proxy.GetMethod("Divide").(func(int, int) (int, error))

	// Use the proxied methods
	fmt.Println("=== Using proxied Add method ===")
	result1 := addMethod(5, 3)
	fmt.Printf("Result: %d\n\n", result1)

	fmt.Println("=== Using proxied Divide method ===")
	result2, err := divideMethod(10, 2)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %d\n", result2)
	}

	fmt.Println("\n=== Using proxied Divide method with error ===")
	result3, err := divideMethod(10, 0)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %d\n", result3)
	}
}
