package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 读写锁demo，常用于读多写少场景
// RWMutex 的零值是未加锁的状态，所以，当你使用 RWMutex 的时候，无论是声明变量，还是嵌入到其它 struct 中，都不必显式地初始化。
// 尽量避免锁重入，重入可能会带来隐蔽的死锁问题

type Counter2 struct {
	mu sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter2) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 使用读锁保护
func (c *Counter2) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestCounter2(t *testing.T) {
	var counter Counter2
	// 10个读
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(counter.Count())
			time.Sleep(time.Millisecond)
		}()
	}
	// 一个写
	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}