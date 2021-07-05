package main

import "fmt"

func main()  {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}

	arry := BubbleSort(arr)
	fmt.Println(arry[len(arr)-1])
}

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
