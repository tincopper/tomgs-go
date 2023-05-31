package main

import (
	"fmt"
)

type Demo struct {
	Name string
}

func (d *Demo) Demoer() string {
	return d.Name
}

func SocketDemo() {
	//con, err := net.Dial("tcp", "127.0.0.1:8080")
	/*func Dial(network, address string) (Conn, error) {
	    // 这个地方很神奇，声明了Dialer类型的变量d然后就可以直接调用方法Dial了
	    var d Dialer
	    return d.Dial(network, address)
	}*/
	// 结构体有默认值，等价于 &TypeName{}
	// var d net.Dialer == var d = net.Dialer{}
	var d Demo
	//d.Name = "123456"
	fmt.Println(d.Demoer())

	var i int
	fmt.Println(i)

	var s string
	fmt.Println(s)
	i2 := len(s)
	fmt.Println(i2)
}

func main1() {
	SocketDemo()
}
