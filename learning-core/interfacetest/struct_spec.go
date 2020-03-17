package main

import (
    "fmt"
)

type Demo struct {
    Name string
}

func (d *Demo) Demoer() string {
    return "123"
}

func SocketDemo() {
    //con, err := net.Dial("tcp", "127.0.0.1:8080")
    /*func Dial(network, address string) (Conn, error) {
        // 这个地方很神奇，声明了Dialer类型的变量d然后就可以直接调用方法Dial了
        var d Dialer
        return d.Dial(network, address)
    }*/
    var d Demo
    fmt.Println(d.Demoer())
}

func main() {
    SocketDemo()
}