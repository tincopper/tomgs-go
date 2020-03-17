package main

import "fmt"

/*
    多个类型可以实现相同的接口
    一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
    也就是说，使用者并不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的。
 */

// 一个服务需要满足能够开启和写日志的功能
type Service interface {
    Start()  // 开启服务
    Log(string)  // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
    fmt.Println(l)
}

// 游戏服务
// 它只实现了start方法，log方法交由Logger进行实现，而把Logger进行组合达到实现log方法的目的，避免了代码的冗余，简化了代码结构
type GameService struct {
    Logger  // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
    fmt.Println("game services start...")
}

func UseService()  {
    var s Service = new(GameService)
    s.Start()
    s.Log("hello")
}