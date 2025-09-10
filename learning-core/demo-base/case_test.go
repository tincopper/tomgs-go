package demo_base

import (
	"fmt"
	"testing"
)

// 测试 switch case 语句
func judge(v int) {
	switch v {
	case 1, 3:
		{
			fmt.Println("v的值为", v)
		}
	default:
		{
			fmt.Println("未匹配到，v的值为", v)
		}
	}
}

func TestCase(t *testing.T) {
	a := 1
	judge(a)

	a = 2
	judge(a)

	a = 3
	judge(a)

}
