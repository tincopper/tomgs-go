package main

import "fmt"

/*
  变量逃逸分析
  go run -gcflags "-m -l" submit_argo.go
  使用 go run 运行程序时，-gcflags 参数是编译参数。其中 -m 表示进行内存分配分析，-l 表示避免程序内联，也就是避免进行程序优化。
 */
func main() {
    // 声明a变量并打印
    var a int
    // 调用空函数
    void()
    //打印a和dummy函数
    fmt.Println(a, dummy(0))
}

func dummy(i int) int {
    var c int
    c = i
    return c
}

func void() {

}
