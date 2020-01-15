package main

import (
    "context"
    "github.com/kataras/iris"
    "time"
    "tomgs-go/learning-iris/base"
)

/**
* @Author: tangzy
* @Date: 2020/1/14 9:15
 */

func main() {
    // iris基本用法
    //app := base.IrisBaseMain()
    // iris MVC
    app := base.IrisMvcMain()

    // 在 Tcp 上监听网络地址 0.0.0.0:8080
    // app.Run(iris.Addr(":8080"))

    // 使用自定义 http.Server进行启动
    // app.Run(iris.Server(&http.Server{Addr: ":8080"}))

    // 使用自定义 net.Listener
    /*l, err := net.Listen("tcp4", ":8080")
    if err != nil {
        panic(err)
    }
    app.Run(iris.Listener(l))*/

    // 优雅关闭
    iris.RegisterOnInterrupt(func() {
        timeout := 5 * time.Second
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        defer cancel()
        // 关闭所有主机
        app.Shutdown(ctx)
    })

    app.Configure()
    app.ConfigureHost()
    app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler, iris.WithCharset("UTF-8"))

}
