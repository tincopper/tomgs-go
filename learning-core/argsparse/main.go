package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

var b = flag.Bool("b", false, "bool类型参数")
var s = flag.String("s", "", "string类型参数")

func main() {
    //for idx, args := range os.Args[1:] {
    for idx, args := range os.Args {
        fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
    }

    // go run main.go -b -s test
    flag.Parse()
    fmt.Println("-b:", *b)
    fmt.Println("-s:", *s)
    fmt.Println("其他参数：", flag.Args())
}
