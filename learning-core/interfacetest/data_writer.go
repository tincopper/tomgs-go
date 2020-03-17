package main

import "fmt"

/*
    接口示例
    1、接口被实现的条件一：接口的方法与实现接口的类型方法格式一致
    2、接口被实现的条件二：接口中所有方法均被实现

    Go语言的接口实现是隐式的，无须让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计。
    实现者在编写方法时，无法预测未来哪些方法会变为接口。一旦某个接口创建出来，要求旧的代码来实现这个接口时，
    就需要修改旧的代码的派生部分，这一般会造成雪崩式的重新编译。

    提示
    传统的派生式接口及类关系构建的模式，让类型间拥有强耦合的父子关系。这种关系一般会以“类派生图”的方式进行。
    经常可以看到大型软件极为复杂的派生树。随着系统的功能不断增加，这棵“派生树”会变得越来越复杂。
    对于 Go语言来说，非侵入式设计让实现者的所有类型均是平行的、组合的。如何组合则留到使用者编译时再确认。
    因此，使用 GO语言时，不需要同时也不可能有“类派生图”，开发者唯一需要关注的就是“我需要什么？”，以及“我能实现什么？”。
 */

// 接口定义
// go接口名称一般以er结尾
type DataWriter interface {
    WriterData(data interface{}) error
}

// 接口实现
// 定义文件结构，用于实现DataWriter
type file struct {

}

// 实现方法，用file结构体进行接受
func (f *file) WriterData(data interface{}) error {
    fmt.Println("WriteData:", data)
    return nil
}

// 使用
func UseDataWriter() {
    // 实例化子类file
    f := new(file)
    // 声明一个接口
    var writer DataWriter
    // 将子类实例赋值为接口
    writer = f
    // 后面就可以通过接口去调用方法了
    err := writer.WriterData("123123")
    if err != nil {
        fmt.Println("error:", err)
    }

    // 还可以这种方式实现
    var s DataWriter = new(file)
    err = s.WriterData("321")
    if err != nil {
        fmt.Println("error:", err)
    }

}