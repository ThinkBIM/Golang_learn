package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newTask goroutine i %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine i %d\n", i)
		time.Sleep(1 * time.Second)
	}

	//main
	//fmt.Println("main goroutine exit")
}
