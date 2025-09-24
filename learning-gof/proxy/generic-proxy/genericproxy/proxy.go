package genericproxy

import (
	"fmt"
	"reflect"
)

// MethodInvocation represents a method call with its arguments
type MethodInvocation struct {
	Method     reflect.Method
	Args       []reflect.Value
	Target     interface{}
	MethodName string
}

// Proceed executes the actual method call
func (mi *MethodInvocation) Proceed() []reflect.Value {
	in := make([]reflect.Value, len(mi.Args)+1)
	in[0] = reflect.ValueOf(mi.Target)
	for i, arg := range mi.Args {
		in[i+1] = arg
	}
	return mi.Method.Func.Call(in)
}

// Interceptor defines the interface for pre/post processing
type Interceptor interface {
	Before(mi *MethodInvocation) error
	After(mi *MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error)
}

// SimpleInterceptor is a basic implementation of Interceptor
type SimpleInterceptor struct {
	BeforeFunc func(mi *MethodInvocation) error
	AfterFunc  func(mi *MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error)
}

func (si *SimpleInterceptor) Before(mi *MethodInvocation) error {
	if si.BeforeFunc != nil {
		return si.BeforeFunc(mi)
	}
	return nil
}

func (si *SimpleInterceptor) After(mi *MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	if si.AfterFunc != nil {
		return si.AfterFunc(mi, result, err)
	}
	return result, err
}

// Proxy holds the target object and interceptors
type Proxy struct {
	target      interface{}
	interceptor Interceptor
}

// NewProxy creates a new proxy for the target object with the given interceptor
func NewProxy(target interface{}, interceptor Interceptor) *Proxy {
	return &Proxy{
		target:      target,
		interceptor: interceptor,
	}
}

// GetMethod returns a function that can be called to invoke the proxied method
func (p *Proxy) GetMethod(methodName string) interface{} {
	targetType := reflect.TypeOf(p.target)
	method, found := targetType.MethodByName(methodName)
	if !found {
		panic(fmt.Sprintf("method %s not found", methodName))
	}

	// Create a function that matches the method signature
	// We need to build a function with the correct number of parameters
	methodType := method.Type

	// Get input types (excluding receiver)
	in := make([]reflect.Type, methodType.NumIn()-1)
	for i := 1; i < methodType.NumIn(); i++ {
		in[i-1] = methodType.In(i)
	}

	// Get output types
	out := make([]reflect.Type, methodType.NumOut())
	for i := 0; i < methodType.NumOut(); i++ {
		out[i] = methodType.Out(i)
	}

	// Create the function type and then the actual function
	funcType := reflect.FuncOf(in, out, methodType.IsVariadic())

	return reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		// Create method invocation
		mi := &MethodInvocation{
			Method:     method,
			Args:       args,
			Target:     p.target,
			MethodName: methodName,
		}

		// Before processing
		if err := p.interceptor.Before(mi); err != nil {
			// Handle error - return zero values for all return types plus the error
			out := make([]reflect.Value, methodType.NumOut())
			for i := 0; i < methodType.NumOut()-1; i++ {
				out[i] = reflect.Zero(methodType.Out(i))
			}
			out[len(out)-1] = reflect.ValueOf(err)
			return out
		}

		// Proceed with the actual method call
		result := mi.Proceed()

		// After processing
		var err error
		// Check if the last return value is an error
		if methodType.NumOut() > 0 && methodType.Out(methodType.NumOut()-1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if !result[len(result)-1].IsNil() {
				err = result[len(result)-1].Interface().(error)
			}
		}

		finalResult, finalErr := p.interceptor.After(mi, result, err)

		// If there was an error in After processing, update the error return value
		if finalErr != nil && methodType.NumOut() > 0 {
			if methodType.Out(methodType.NumOut() - 1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
				finalResult[len(finalResult)-1] = reflect.ValueOf(finalErr)
			}
		}

		return finalResult
	}).Interface()
}
