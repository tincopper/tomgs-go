package demo3_struct

import "fmt"

// 代理函数，实现方法与函数的统一调用
type Class struct {

}

// 给结构体添加Do方法
func (c *Class) Do(v int)  {
    fmt.Println("call method do：", v)
}

// 普通函数
func funcDo(v int)  {
    fmt.Println("call function do:", v)
}

// ------------------------------------------------------
func UseDelegate()  {
    // 声明一个函数回调
    var delegate func(int)

    c := new(Class)
    delegate = c.Do
    delegate(1)
    // 将回调设为普通函数
    delegate = funcDo
    delegate(2)
}