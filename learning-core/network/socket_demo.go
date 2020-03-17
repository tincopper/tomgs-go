package network

import (
    "net"
)
func SocketDemo() {
    con, err := net.Dial("tcp", "127.0.0.1:8080")
    
}

func main() {
    SocketDemo()
}