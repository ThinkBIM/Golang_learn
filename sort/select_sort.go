package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100, 100)

	for i := 0; i < 100; i++ {
		arr[i] = i * rand.Intn(100)
	}


	fmt.Println(SelectMax(arr))
	fmt.Println(SelectSort(arr))
}

func SelectMax(arr []int) int {
	if len(arr) <= 1 {
		return arr[0]
	}

	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	for i := 1; i < length; i++ {
		min := i
		for j := i+1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr

}

