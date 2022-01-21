package main

import (
	"fmt"
	"sync"
	"testing"
)

// Once 可以用来执行且仅仅执行一次动作，常常用于单例对象的初始化场景。
// 主要用于延迟初始化的场景
// 并发访问只需初始化一次的共享资源

func TestOnce(t *testing.T) {
	var once sync.Once

	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1)

	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2) // 无输出
}