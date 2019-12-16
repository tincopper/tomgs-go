package demo1

import (
    "errors"
    "fmt"
    "runtime"
)

// 定义除数为0的情况并返回
var errorDivisionByZero = errors.New("division by zero")

// 除法
func Div(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errorDivisionByZero
    }

    return dividend /divisor, nil
}

// 宕机
func Panic()  {
    defer fmt.Println("宕机后要做的事情1")
    defer fmt.Println("宕机后要做的事情2")
    panic("crash")
}

type panicContext struct {
    function string // 所在函数
}

func ProtectRun(entry func()) {
    // 定义延迟函数用于catch异常
    defer func() {
        // 发生宕机时进行异常捕获
        err := recover()
        switch err.(type) {
        case runtime.Error: // 运行时异常
            fmt.Println("runtime error:", err)
        default: // 非运行时异常
            fmt.Println("error:", err)
        }
    }()
    // 执行业务逻辑
    entry()
}

// 异常catch模拟
func RecoverMock() {
    fmt.Println("运行前...")

    // 模拟手动宕机
    ProtectRun(func() {
        fmt.Println("手动宕机前")
        //使用panic传递上下文
        panic(&panicContext {
            "手动触发panic",
        })
        fmt.Println("手动宕机后")
    })
    // 模拟空指针异常
    ProtectRun(func() {
        fmt.Println("赋值宕机前")
        var a *int
        *a = 1
        fmt.Println("赋值宕机后")
    })
    fmt.Println("运行后...")
}