package func_demo

import (
    "fmt"
    "testing"
)

func TestDD(t *testing.T) {
    ee := func(x, y int) int {
        return x + y
    }
    result := Dd(ee)
    fmt.Println(result)
}