package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"testing"
)

func TestUseChannel(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(debug.Stack()))
			return
		}
	}()
	UseChannel()
}

func TestUseChannel2(t *testing.T) {
	UseChannel2()
}

func Count(ch chan int) {
	fmt.Println("counting ...")
	ch <- 1
}

func TestCount(t *testing.T) {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		result := <-ch
		fmt.Println(result)
	}
}
