package demo_base

import (
    "fmt"
    "strconv"
)

// 整型转字符
func IntToStr(a int) string {
    return strconv.Itoa(a)
}

// 字符串转整型
func StrToInt(str string) int {
    result, err := strconv.Atoi(str)
    if err != nil {
        fmt.Printf("%s 转换失败", str)
        panic(err)
    }
    return result
}

// 其他的转换函数
func ParseOther()  {
    //strconv.ParseBool()
    //strconv.ParseFloat()
    //strconv.ParseInt()
    //strconv.ParseUint()
}