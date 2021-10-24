package demo_base

import (
	"fmt"
	"testing"
)

/**
常量使用一个名称来绑定一块内存地址，该内存地址中存放的数据类型由定义常量时指定
的类型决定，而且该内存地址里面存放的内容不可以改变 。 Go 中常量分为布尔型、宇符串型和
数值型常量。常量存储在程序的只读段里（ .rodata section ） 。

预声明标识符 iota 用在常量声明中，其初始值为 0。一组多个常量同时声明时其值逐行增
加， iota 可以看作自增的枚举变量，专 门用来初始化常量。
*/
func TestIota(t *testing.T) {
	//类似枚举的 iota
	const (
		cO = iota //cO == 0
		c1 = iota //cl == 1
		c2 = iota //c2 == 2
	)
	fmt.Println(cO)
	fmt.Println(c1)
	fmt.Println(c2)

	// 简写模式
	const (
		c3 = iota
		c4
		c5
	)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
}

func TestIota2(t *testing.T) {
	const (
		a = 1 << iota // a == 1 (iota == 0)
		b = 1 << iota // b == 2 (iota == 1)
		c = 3         // b == 3 (iota == 2, iota unused)
		d = 1 << iota // b == 8 (iota == 3)
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func TestIota3(t *testing.T) {
	const (
		u         = iota * 42   // u = 0 (untyped integer constant)
		v float64 = iota * 42.0 // v == 42.0 (float64 constant)
		w         = iota * 42   // w == 84 (untyped integer constant)
	)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)

	// 分开的 const 语句， iota 每次都从 0 开始
	const x = iota // x == 0
	const y = iota // y == 0
	fmt.Println(x)
	fmt.Println(y)
}
