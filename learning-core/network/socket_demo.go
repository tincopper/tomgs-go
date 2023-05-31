package network

import (
	"fmt"
	"net"
)

func SocketDemo() {
	con, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Printf("con.LocalAddr().Network(): %v\n", con.LocalAddr().String())
}

func main() {
	SocketDemo()
}
