package main

import (
	"fmt"
	"sync"
	"testing"
)

// base using
func TestMutex(t *testing.T) {
	mutex := sync.Mutex{}
	// or
	//var mutex sync.Mutex
	mutex.Lock()
	// do something ...
	defer mutex.Unlock()
}

// data race demo
// 10个协程对count进行10w次的累加
func TestMutex2(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//  count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// data race demo2
func TestMutex3(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//  count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// mutex optimized
func TestMutex4(t *testing.T) {
	counter := &Counter{Name: "count demo"}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//  count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 在业务中使用Mutex的更好的姿势
type Counter struct {
	// 不需要包护的资源
	Name string

	// mutex 放在被保护资源的上面
	mu sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
