package main

import (
	"fmt"
	"reflect"
	"time"
	"tomgs-go/learning-gof/proxy/generic-proxy/genericproxy"
)

// UserService defines operations for user management
type UserService interface {
	GetUser(id int) (string, error)
	CreateUser(name string, age int) (int, error)
}

// userService is the implementation of UserService
type userService struct {
	users  map[int]string
	nextID int
}

func NewUserService() UserService {
	return &userService{
		users:  make(map[int]string),
		nextID: 1,
	}
}

func (s *userService) GetUser(id int) (string, error) {
	if name, exists := s.users[id]; exists {
		return name, nil
	}
	return "", fmt.Errorf("user with id %d not found", id)
}

func (s *userService) CreateUser(name string, age int) (int, error) {
	id := s.nextID
	s.nextID++
	s.users[id] = name
	fmt.Printf("Created user: %s (ID: %d, Age: %d)\n", name, id, age)
	return id, nil
}

// TimingInterceptor measures execution time of methods
type TimingInterceptor struct {
	startTime time.Time
}

func (ti *TimingInterceptor) Before(mi *genericproxy.MethodInvocation) error {
	ti.startTime = time.Now()
	fmt.Printf("Starting execution of %s at %v\n", mi.MethodName, ti.startTime)
	return nil
}

func (ti *TimingInterceptor) After(mi *genericproxy.MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	duration := time.Since(ti.startTime)
	fmt.Printf("Finished execution of %s in %v\n", mi.MethodName, duration)
	return result, err
}

// ValidationInterceptor validates input parameters
type ValidationInterceptor struct{}

func (vi *ValidationInterceptor) Before(mi *genericproxy.MethodInvocation) error {
	// Validate that ID is positive for GetUser
	if mi.MethodName == "GetUser" && len(mi.Args) > 0 {
		id := mi.Args[0].Int()
		if id <= 0 {
			return fmt.Errorf("invalid user ID: %d", id)
		}
	}

	// Validate that age is reasonable for CreateUser
	if mi.MethodName == "CreateUser" && len(mi.Args) > 1 {
		age := mi.Args[1].Int()
		if age < 0 || age > 150 {
			return fmt.Errorf("invalid age: %d", age)
		}
	}

	return nil
}

func (vi *ValidationInterceptor) After(mi *genericproxy.MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	return result, err
}

// ChainedInterceptor combines multiple interceptors
type ChainedInterceptor struct {
	interceptors []genericproxy.Interceptor
}

func NewChainedInterceptor(interceptors ...genericproxy.Interceptor) *ChainedInterceptor {
	return &ChainedInterceptor{
		interceptors: interceptors,
	}
}

func (ci *ChainedInterceptor) Before(mi *genericproxy.MethodInvocation) error {
	for _, interceptor := range ci.interceptors {
		if err := interceptor.Before(mi); err != nil {
			return err
		}
	}
	return nil
}

func (ci *ChainedInterceptor) After(mi *genericproxy.MethodInvocation, result []reflect.Value, err error) ([]reflect.Value, error) {
	finalResult := result
	finalErr := err

	for i := len(ci.interceptors) - 1; i >= 0; i-- {
		finalResult, finalErr = ci.interceptors[i].After(mi, finalResult, finalErr)
	}

	return finalResult, finalErr
}

func main() {
	// Create the target service
	userService := NewUserService()

	// Create interceptors
	timingInterceptor := &TimingInterceptor{}
	validationInterceptor := &ValidationInterceptor{}

	// Create a chained interceptor
	chainedInterceptor := NewChainedInterceptor(validationInterceptor, timingInterceptor)

	// Create the proxy
	proxy := genericproxy.NewProxy(userService, chainedInterceptor)

	// Get proxied methods
	createUserMethod := proxy.GetMethod("CreateUser").(func(string, int) (int, error))
	getUserMethod := proxy.GetMethod("GetUser").(func(int) (string, error))

	// Use the proxied methods
	fmt.Println("=== Creating a user ===")
	id, err := createUserMethod("Alice", 25)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
	} else {
		fmt.Printf("Created user with ID: %d\n\n", id)
	}

	fmt.Println("=== Getting a user ===")
	name, err := getUserMethod(id)
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
	} else {
		fmt.Printf("Found user: %s\n\n", name)
	}

	fmt.Println("=== Trying to create a user with invalid age ===")
	_, err = createUserMethod("Bob", -5)
	if err != nil {
		fmt.Printf("Expected error: %v\n\n", err)
	}

	fmt.Println("=== Trying to get a non-existent user ===")
	_, err = getUserMethod(999)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}
