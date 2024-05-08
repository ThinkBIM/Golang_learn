package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// Onces 只会执行一次
func Onces() {
	once.Do(func() {
		fmt.Println("执行一次")
	})
}

func main() {
	Onces()
}
