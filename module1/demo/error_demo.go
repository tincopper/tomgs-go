package demo

import "errors"

// 定义除数为0的情况并返回
var errorDivisionByZero = errors.New("division by zero")

func Div(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errorDivisionByZero
    }

    return dividend /divisor, nil
}