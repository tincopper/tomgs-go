package slice

import (
    "fmt"
    "testing"
)

// 切片赋值
func TestSliceAssign(t *testing.T) {
    // 方式一 ：由数组定义slice
    var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var slice = array[0:2]
    fmt.Printf("%d\n", slice)
    
    //方式二：直接创建数组切片  然后一个一个的赋值
    slice2 := make([]int, 5, 10)
    slice2[0] = 100
    fmt.Printf("%d\n", slice2)
    
    //方式三：使用append进行赋值
    slice0 := make([]int, 5, 10)
    slice0 = append(slice0, 444, 555)
    fmt.Printf("%d\n", slice0)
    
    //方式四：直接创建并初始化数组切片
    slice3 := []int{1, 2, 3, 4, 5, 6}
    fmt.Printf("%d\n", slice3)
    
    //方式五.基于数组切片的进行赋值
    slice4 := slice3[:4]
    fmt.Printf("%d\n", slice4)
}

func TestSliceExpand(t *testing.T) {
    fmt.Println("--------")
}

// 4、基于数组切片创建数组切片
// 类似于数组切片可以基于一个数组创建，数组切片也可以基于另一个数组切片创建。
func TestNewSliceOnSlice(t *testing.T) {
    oldSlice := []int{1, 2, 3, 4, 5}
    newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片
    fmt.Println("newSlice:", newSlice)
    
    /*
       有意思的是，选择的oldSlice元素范围甚至可以超过所包含的元素个数，比如newSlice
       可以基于oldSlice的前6个元素创建，虽然oldSlice只包含5个元素。只要这个选择的范围不超
       过oldSlice存储能力（即cap()返回的值），那么这个创建程序就是合法的。 newSlice中超出
       oldSlice元素的部分都会填上0。
    */
    oldSlice1 := make([]int, 0, 10)
    capV := cap(oldSlice1)
    fmt.Println("cap:", capV)
    
    // 切片赋值
    oldSlice1 = append(oldSlice1, 1, 2, 3)
    fmt.Println("oldSlice1:", oldSlice1)
    
    newSlice1 := oldSlice1[:6]
    fmt.Println("newSlice1:", newSlice1)
}

// 5、切片复制
// 数组切片支持Go语言的另一个内置函数copy()，用于将内容从一个数组切片复制到另一个
// 数组切片。如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制
func TestSliceCopy(t *testing.T) {
    slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{5, 4, 3}
    copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
    fmt.Println("slice1:", slice1, "slice2:", slice2)
    copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
    fmt.Println("slice1:", slice1, "slice2:", slice2)
}
