package main

import (
	"fmt"
	"tomgs-go/module1/demo"
)

func main() {
	// 基础的demo练习
	baseDemo()
	// 类型转换
	demo.TypeConvert()
	// 指针
	demo.GoPointer()

	// 通过下面两个示例可以说明go也是值传递
	x, y := 1, 2
	// 交换变量值
	swapInt(1, 2)
	fmt.Printf("x = %d, y = %d\n", x, y)

	swapInt2(&x, &y)
	fmt.Printf("x = %d, y = %d\n", x, y)
	//
	command := demo.ParseCommand()
	fmt.Println(command)
	//
	intToStr := demo.IntToStr(1)
	fmt.Printf("%T, %#v\n", intToStr, intToStr)

	strToInt := demo.StrToInt("100")
	fmt.Printf("%T, %#v\n", strToInt, strToInt)

	//strToInt2 := demo.StrToInt("s100")
	//fmt.Printf("%T, %#v\n", strToInt2, strToInt2)

	// 切片 a[1:2] 左闭右开区间，取出的元素数量为：结束位置 - 开始位置；
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	fmt.Println(a, a[1:])
	fmt.Println(a, a[:2])
	fmt.Println(a, a[:])
	fmt.Println(a, a[0:0])
	fmt.Println(a, a[:0])

	demo.MapSimple()

	fmt.Println(demo.Div(1, 0))
}

func baseDemo() {
	var a int
	var b int8
	var c = 100
	var str string
	var str1 string = "123123"
	var str2 = "123123"

	a = 200
	b = 127
	i := 200
	str = "123321"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(i)
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println("hello world" + "123")
}

func swapInt(a int, b int) {
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%d%s%d\n", a, "--", b)
}

// debug可以发现，传入的是地址，地址是没有变的，但是地址指向的值发生了改变
func swapInt2(a, b *int) {
	// 取a指针的值赋值给tmp
	tmp := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = tmp
}