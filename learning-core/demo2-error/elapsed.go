package demo2_error

import (
    "fmt"
    "time"
)

// 方法运行耗时统计
func MethodElapsed()  {
    start := time.Now()
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum++
    }
    elapsed := time.Since(start)
    fmt.Println("cost time:", elapsed)
}