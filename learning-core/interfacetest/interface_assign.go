package main

// 接口赋值
// 将对象实例赋值给接口；
// 将一个接口赋值给另一个接口；
type Integer int

// 实现接口方法的两种写法
func (a Integer) Less(b Integer) bool {
	return a < b
}

/*func (a Integer) Add(b Integer) {
	a += b
}*/

func (a *Integer) Add(b Integer) {
	*a += b
}

// --------------------------------------------------
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

// --------------------------------------------------
type Lesser interface {
	Less(b Integer) bool
}

func main2() {
	// 将a 赋值给接口LessAdder
	//var a Integer = 1
	// 这样赋值是错误的
	// Cannot use 'a' (type Integer) as type LessAdder in assignment Type does not implement 'LessAdder'
	// as 'Add' method has a pointer receiver
	//var b LessAdder = a
	// 正确赋值
	//var c LessAdder = &a

	// 将a 赋值给接口Lesser
	//var a Integer = 1
	// 正确赋值
	//var b1 Lesser = &a
	// 正确赋值
	//var b2 Lesser = a
}
