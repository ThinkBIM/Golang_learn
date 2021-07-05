package main

import "fmt"

func main() {
	arr := []int{1,10,20,2,345,6,21}
	fmt.Println(QuickSort(arr))
}


func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	current := arr[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, current)

	for i:=1;i<len(arr);i++ {
		if arr[i] < current {
			low = append(low, arr[i])
		}else if arr[i] > current {
			high = append(high, arr[i])
		}else {
			mid = append(mid, arr[i])
		}
	}

	low, high = QuickSort(low), QuickSort(high)

	marry := append(append(low, mid...), high...)

	return marry


}
