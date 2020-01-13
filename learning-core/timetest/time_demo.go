package timetest

import (
    "fmt"
    "time"
)

// 时间格式化，一定要是2006-01-02 15:04:05这个时间，否则会出现时间紊乱
func TimeFormat() {
    format := time.Now().Format("20060102150405")
    fmt.Println(format)

    format1 := time.Now().Format("2006-01-02 15:04:05")
    fmt.Println(format1)

    format2 := time.Now().Format("2007-01-02 15:04:05")
    fmt.Println(format2)
}

func TimerTest() {
    timer := time.NewTimer(3 * time.Second)
    fmt.Println("==========")
    <- timer.C

    go func() {
        ticker := time.NewTicker(time.Second * 3)
        for range ticker.C {
            fmt.Printf("ticked at %v\n", time.Now())
        }
    }()
}