package demo3_struct

import "fmt"

type Cat struct {
    Color string
    Name string
}

// ------------------------------------------------------------

// 初始化结构体（模拟构造函数）
func NewCatByName(name string) *Cat {
    return &Cat{
        Name: name,
    }
}

// 初始化结构体
func NewCatByColor(color string) *Cat {
    return &Cat{
        Color: color,
    }
}

// ------------------------------------------------------------
// 模拟子类
type BlackCat struct {
    Cat // 嵌入Cat, 类似于派生，可以进行结构体的嵌入
}

// 构造基类
func NewCat(name string) *Cat {
    return &Cat{
        Name: name,
    }
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
    cat := &BlackCat{}
    cat.Color = color
    return cat
}

type WhiteCat struct {
    // 嵌入Cat, 类似于派生，可以进行结构体的嵌入
    Cat
    // 定义子类自己的属性
    Age int
}

// “构造子类”
func NewWhiteCat(name string, age int) *WhiteCat {
    cat := &WhiteCat{}
    cat.Age = age
    cat.Name = name
    cat.Color = "white"
    return cat
}

// -----------------------------------------------------------------------------
// 结构体方法
func (cat *Cat) SetName(name string)  {
    cat.Name = name
}

func (cat *Cat) SetColor(color string)  {
    cat.Color = color
}

func (cat *Cat) GetName() string {
    return cat.Name
}

func (cat *Cat) GetColor() string {
    return cat.Color
}

// -----------------------------------------------------------------------------
// 使用结构体
func UseStruct()  {
    // 实例化结构体
    // 1、通过new的方式
    whiteCat := new(WhiteCat)
    whiteCat.Name = "tomgs"
    whiteCat.Color = "white"
    whiteCat.Age = 18
    fmt.Println(whiteCat)

    // 2、通过声明的方式
    var blackCat BlackCat
    blackCat.Name = "tomgs"
    blackCat.Color = "black"
    fmt.Println(blackCat)

    // 3、通过取址符
    var blackCat1 = &BlackCat{}
    blackCat1.Name = "tomgs1"
    blackCat1.Color = "black"
    fmt.Println(blackCat1)

    // 4、直接赋值
    blackCat = BlackCat{Cat{
        Color: "black",
        Name:  "tomgs2",
    }}
    fmt.Println(blackCat1)

    // 使用方法
    cat := new(Cat)
    cat.SetName("tomgs")
    cat.SetColor("white")
    fmt.Println(cat)

    name := cat.GetName()
    color := cat.GetColor()
    fmt.Println(name, color)
    
    // 特殊使用
    //s := struct{}{}
    //s1 := struct{a string}{"a"}
}