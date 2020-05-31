package main

import "fmt"

// 指针
// go里面不支持指针的运算
func main() {
	testBasePoint()
	testPointArray()
}

func testPointArray() {
	// 指针数组
	// 首先是一个数组，然后数组里面存放的值是一个指针地址
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组 pointArr:", pointArr)

	// 数组指针
	// 首先是一个指针，然后指针指向的是一个数组
	arr := [...]int{3, 4, 5}
	arrPoint := &arr
	fmt.Println("数组指针 arrPoint:", arrPoint)
}

func testBasePoint() {
	var count int = 30
	// 声明一个指针
	var countPoint *int
	// 指针赋值
	countPoint = &count

	fmt.Printf("count 的地址 %x\n", &count)
	fmt.Printf("countPoint 的地址 %x\n", countPoint)
	fmt.Printf("countPoint 的地址的值 %d\n", *countPoint)
}
