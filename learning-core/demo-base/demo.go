/**
* @Author: tangzy
* @Date: 2019/12/14 13:54
 */
package demo_base

import (
    "flag"
    "fmt"
)

// 类型转换
// Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明
// 当从类型范围较大的转换为范围类型较小的时候会出现精度丢失，这个需要注意一下
// 只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型），不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）
func TypeConvert() {
    // 将浮点类型转为int类型
    i := 5.0
    b := int(i)
    fmt.Println(b)
}

// go指针
// Go语言为程序员提供了控制数据结构指针的能力，但是，并不能进行指针运算。
// 1、类型指针：允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
// 2、切片：由指向起始元素的原始指针、元素数量和容量组成。
// &操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值
/*
    取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

    变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
    1、对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
    2、指针变量的值是指针地址。
    3、对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。
 */
func GoPointer() {
    var i int = 1
    var str string = "abc"
    fmt.Println(&i, &str)
    // %p： 打印变量的内存地址
    fmt.Printf("%p, %p\n", &i, &str)

    // 取地址操作
    var str2 string = "pointer simple"
    // 取地址操作，ptr的类型为*string
    ptr := &str2
    // %T 打印ptr的类型
    fmt.Printf("ptr type : %T\n", ptr)
    fmt.Printf("ptr addreess: %p\n", ptr)

    // 取值操作
    value := *ptr
    fmt.Printf("value type : %T\n", value)
    fmt.Printf("value: %s\n", value)

    // 不能对普通变量进行取值操作
    //s := *str2
    //fmt.Printf("value type : %T\n", s)

    // 对指针取地址，i2类型为**string，双重指针
    i2 := &ptr
    fmt.Printf("value type : %T\n", i2)
    fmt.Println(i2)

    // new的方式创建指针
    str3 := new(string)
    *str3 = "12312"
    fmt.Println(*str3)
}

//定义命令行
var mode = flag.String("mode", "", "process mode")

// 解析命令行
func ParseCommand() string {
    // 解析命令行参数
    flag.Parse()
    // 输出命令行参数
    fmt.Println(*mode)
    return *mode
}
