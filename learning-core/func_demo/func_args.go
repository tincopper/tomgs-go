package func_demo

// 函数式参数
func Dd(i func(int, int) int) int {
    return i(1, 2)
}