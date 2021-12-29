package main

import (
    "fmt"
    "testing"
    "time"
)

/*
    go并发
    eg：go 函数名( 参数列表 )

    使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略。
    如果需要在 goroutine 中返回数据，请使用后面介绍的通道（channel）特性，通过通道把数据从 goroutine 中作为返回值传出。
 */

func running() {
    var times int
    for {
       times++
       fmt.Println("tick ", times)
       // sleep 1s
       time.Sleep(time.Second)
    }
}

// 不要通过共享内存来通信，而应该通过通信来共享内存
//
// channel是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声
// 明channel时指定。如果对Unix管道有所了解的话，就不难理解channel，可以将其认为是一种类
// 型安全的管道。
//
// 经常会遇到需要实现条件等待的场景，这也是channel可以发挥作用的地方。
// 在go中如果遇到需要阻塞，加锁的地方先考虑能否使用channel来处理
func TestRunningGoroutine(t *testing.T) {
    // 开启并发执行
    go running()

    // 测试channel
    UseChannel2()

    // 接受用户输入
    var input string
    _, err := fmt.Scanln(&input)
    if err != nil {
        fmt.Println("error: ", err)
    }
    fmt.Println(input)
}
