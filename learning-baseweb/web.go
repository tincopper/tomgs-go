package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "tomgs-go/learning-baseweb/web"
    "tomgs-go/learning-baseweb/web/controller"
)

/**
* @Author: tangzy
* @Date: 2019/12/17 10:43
 */
func main() {
    //web.StartServerDemo()
    go func() {
        /*web.SetFirstCallBack(func(writer http.ResponseWriter, request *http.Request) bool {
            return true
        })
        err := web.Start(":9090")*/

        err := web.LoadRouteAndStartServe(func() {
            controller.LoadControllers()
        }, ":9090")

        if err != nil {
            // 记录日志并退出
            log.Fatal("start server error:", err)
        }
    }()

    listenSignal()
}

func init() {
    fmt.Println("-------------------")
    controller.LoadControllers()
}

func listenSignal() {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
    for sig := range signals {
        if sig == nil {
            continue
        }
        //server.StopAllChannel()
        //os.Remove(*BifrostPid)
        //server.Close()
        os.Exit(0)
    }
}
