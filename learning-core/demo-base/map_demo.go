package demo_base

import "fmt"

/*
    集合：map操作
 */
func MapSimple()  {
    // 声明一个map集合
    var mapLit map[string] int
    // 分配内存，如果没有下面的操作直接给mapList添加元素会出现nil
    // mapLit = map[string]int{"one": 1, "two": 2}
    mapLit = map[string] int{}
    // 创建map集合
    mp1 := make(map[string] int)
    mp2 := make(map[string] int, 100) // 指定容量
    // 用切片作为 map 的值
    mp3 := make(map[int] []int)
    mp4 := make(map[int] *[]int)
    // 赋值
    mapLit["aa"] = 1
    mapLit["bb"] = 2
    mp1["a"] = 1
    mp1["b"] = 2
    mp2["a"] = 1
    mp2["b"] = 2
    mp3[1] = []int{1, 2, 3}
    mp4[1] = &[]int{1, 2, 3}
    // 遍历
    // 1、获取指定的key的值
    i := mapLit["aa"]
    fmt.Println(i)

    // 2、循环遍历

    // 删除

    fmt.Println(mapLit, mp1, mp2)
}