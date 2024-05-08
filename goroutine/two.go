package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("defer A")
		fmt.Println("A")

		func() {
			defer fmt.Println("defer B")
			runtime.Goexit() //退出当前goroutine
			fmt.Println("B")
		}()
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
