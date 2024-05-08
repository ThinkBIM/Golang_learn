package main

import (
	"fmt"
	"sync"
)

func main() {
	res := 0
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(val int) {
			res += val
			wg.Done()
		}(i)
	}

	// 这里什么结果都可能拿到
	// wg.Wait()
	fmt.Println(res)
}
