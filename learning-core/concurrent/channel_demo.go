package main

import (
    "fmt"
    "time"
)

func UseChannel() {
    // 声明通道，要配合make使用，但是直接使用make就可以了
    var ch chan int
    // 创建通道
    ch1 := make(chan int)
    ch = ch1

    // 往通道写入数据，如果写入的数据没有接受的话会抛出fatal error: all goroutines are asleep - deadlock!
    ch <- 1
    ch1 <- 0
}

func UseChannel2() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 3; i++ {
            // 往通道循环写入数据
            ch <- i
            // 睡眠1s
            time.Sleep(time.Second)
        }
    }()
    // 循环接受数据
    /*for {
        data := <- ch
        fmt.Println(data)
        if data == 2 {
            break
        }
    }*/
    go func() {
        for data := range ch {
            fmt.Println(data)
            if data == 2 {
                break
            }
        }
    }()
}
