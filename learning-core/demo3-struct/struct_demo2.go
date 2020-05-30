package demo3_struct

import "fmt"

type Point struct {
    X int
    Y int
}

// 非指针类型接收器的加方法
func (p Point) Add(point Point) Point {
    // 相加后返回一个新的Point对象
    return Point{
        X: p.X + point.X,
        Y: p.Y + point.Y,
    }
}

// 小对象由于值复制时的速度较快，所以适合使用非指针接收器，大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。
func UsePoint()  {
    point1 := Point{1, 1}
    point2 := Point{2, 2}

    point3 := point1.Add(point2)
    fmt.Println(point3)
}



