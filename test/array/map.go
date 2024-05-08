package main

import (
	"fmt"
	"sort"
)

func main()  {
	items := make([]map[int]int, 5)
	for i:= range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
		items[i][2] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	//	Version A: Value of items: [map[1:2 2:2] map[1:2 2:2] map[1:2 2:2] map[1:2 2:2] map[1:2 2:2]]

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
	//Version B: Value of items: [map[] map[] map[] map[] map[]]


	//sort
	sorts()
}

func sorts()  {
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)

	inmap := make(map[int]string, len(barVal))

	for s, i := range barVal {
		inmap[i] = s
	}

	fmt.Println(inmap)

	fmt.Println("unsorted:")
	fmt.Println(barVal)
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}