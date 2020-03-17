package main

import "fmt"

func Count(ch chan int) {
    fmt.Println("counting ...")
    ch <- 1
}

func main() {
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