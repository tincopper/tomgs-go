package demo3_struct

import "fmt"

/*
  自定义类型：
  在Go语言中，使用 type 关键字可以定义出新的自定义类型，之后就可以为自定义类型添加各种方法了
 */

// 自定义类型
type MyInt int

// 是否为0
func (myInt MyInt) IsZero() bool {
    return myInt == 0
}

// add方法
func (myInt MyInt) Add(i int) int {
    return int(myInt) + i
}

// -------------------------------------------------
func UseMyInt()  {
    var i MyInt
    zero := i.IsZero()
    fmt.Println(zero)

    i = 0
    result := i.Add(2)
    fmt.Println(result)
}
