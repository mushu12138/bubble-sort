package main

import "fmt"

// BubbleSort 对整型切片进行冒泡排序（升序），原地排序并返回该切片。
func BubbleSort(arr []int) []int {
	n := len(arr)
	// 外层控制轮数，每轮把当前最大值“冒泡”到末尾
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// 整轮无交换说明已有序，提前结束
		if !swapped {
			break
		}
	}
	return arr
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前:", data)
	BubbleSort(data)
	fmt.Println("排序后:", data)
}
