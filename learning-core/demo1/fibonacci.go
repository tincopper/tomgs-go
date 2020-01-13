package demo1

import (
    "fmt"
    "time"
)

const LIMIT int = 40
var fibs [LIMIT]int

// 斐波拉契
// 1 1 2 3 5 8 13
// 计算数列中第 n 个数字
func fibonacci(n int) (res int) {
    if n <= 2 {
        res = 1
    }
    res = fibonacci(n - 1) + fibonacci(n - 2)
    return
}

func FibMock() {
    result := 0
    start := time.Now()
    for i := 1; i < LIMIT; i++ {
        //result = fibonacci(i)
        result = fibonacci2(i)
        fmt.Printf("数列第 %d 位: %d\n", i, result)
    }
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("程序的执行时间为: %s\n", delta)
}

func fibonacci2(n int) (res int) {
    if fibs[n] != 0 {
        res = fibs[n]
        return
    }
    if n <= 2 {
        res = 1
    } else {
        res = fibonacci2(n - 1) + fibonacci2(n - 2)
    }

    fibs[n] = res
    return
}