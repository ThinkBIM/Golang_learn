package main

import "fmt"

func main() {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024 * 1024; i++ {
		arr[i] = i+1
	}

	id := binSearch(arr, 10)

	if id != -1 {
		fmt.Println(id, arr[id])
	}else {
		fmt.Println("没有找到数据")
	}

	//
	//fmt.Println(arr)
}

func binSearch(arr []int, finddata int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if arr[mid] > finddata {
			high = mid - 1
		} else if arr[mid] < finddata {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}