package genericproxy

import (
	"errors"
	"reflect"
	"testing"
)

// mockService is a simple service for testing
type mockService struct {
	value int
}

func (s *mockService) GetValue() int {
	return s.value
}

func (s *mockService) SetValue(v int) {
	s.value = v
}

func (s *mockService) ErrorMethod() error {
	return errors.New("test error")
}

// testInterceptor is a simple interceptor for testing
type testInterceptor struct {
	beforeCalled bool
	afterCalled  bool
}

func (ti *testInterceptor) Before(mi *MethodInvocation) error {
	ti.beforeCalled = true
	return nil
}

func (ti *testInterceptor) After(mi *MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	ti.afterCalled = true
	return result, err
}

func TestProxyGetMethod(t *testing.T) {
	service := &mockService{value: 42}
	interceptor := &testInterceptor{}

	proxy := NewProxy(service, interceptor)

	// Test GetValue method
	getValueMethod := proxy.GetMethod("GetValue").(func() int)
	result := getValueMethod()

	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}

	if !interceptor.beforeCalled {
		t.Error("Before method was not called")
	}

	if !interceptor.afterCalled {
		t.Error("After method was not called")
	}
}

func TestProxySetMethod(t *testing.T) {
	service := &mockService{value: 0}
	interceptor := &testInterceptor{}
	interceptor.beforeCalled = false
	interceptor.afterCalled = false

	proxy := NewProxy(service, interceptor)

	// Test SetValue method
	setValueMethod := proxy.GetMethod("SetValue").(func(int))
	setValueMethod(100)

	if service.value != 100 {
		t.Errorf("Expected 100, got %d", service.value)
	}

	if !interceptor.beforeCalled {
		t.Error("Before method was not called")
	}

	if !interceptor.afterCalled {
		t.Error("After method was not called")
	}
}

func TestProxyErrorMethod(t *testing.T) {
	service := &mockService{}
	interceptor := &testInterceptor{}
	interceptor.beforeCalled = false
	interceptor.afterCalled = false

	proxy := NewProxy(service, interceptor)

	// Test ErrorMethod
	errorMethod := proxy.GetMethod("ErrorMethod").(func() error)
	err := errorMethod()

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if err.Error() != "test error" {
		t.Errorf("Expected 'test error', got '%s'", err.Error())
	}

	if !interceptor.beforeCalled {
		t.Error("Before method was not called")
	}

	if !interceptor.afterCalled {
		t.Error("After method was not called")
	}
}
