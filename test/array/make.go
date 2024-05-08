package main

import (
	"fmt"
	"reflect"
)

var arr [5]int

func main()  {
	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		arr[i] = i * 2
	}

	for key, value := range arr {
		fmt.Println("Array key", "is", key)
		fmt.Println("Array value", "is", value)
	}

	for i := range arr {
		fmt.Println("Array key", "is", i, arr[i])
	}

	fmt.Println(arr)

	t := reflect.TypeOf(arr)
	fmt.Println(t)

	var arrLazy = []int{5, 6, 7, 8, 22}
	fmt.Println(arrLazy)

	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)
}