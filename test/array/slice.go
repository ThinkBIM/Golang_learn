package main

import "fmt"

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(len(ar[5:6]))
	fmt.Println(len(ar[5:5]))

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

}