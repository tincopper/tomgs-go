package main

import (
	"fmt"
	"sync"
	"time"
)

var WG sync.WaitGroup

func main() {
	// 协程同步
	Read()
	go Write()
	WG.Wait()
	fmt.Println("All done")
}

func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Done ", i)
		WG.Done()
	}
}

func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}
