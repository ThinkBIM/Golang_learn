package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("cat", "modd")
	m.Store("dog", "wight")

	val, ok := m.Load("cat")
	if ok {
		fmt.Println(ok)
	}

	t, ok := val.(string)
	println(t, ok)
}
