package main

import "fmt"

// 使用二分法查找预期的数值i
// list为有序数组，这个是用二分查找的必要条件
func BinarySearch(list []int, i int) bool {
	low := 0
	high := len(list) - 1

	for low <= high {
		// 找到中间值
		mid := (low + high) / 2
		if mid == i {
			return true
		}
		// 小了
		if mid < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(BinarySearch(list, 2)) // true
	fmt.Println(BinarySearch(list, -1)) // false
}
