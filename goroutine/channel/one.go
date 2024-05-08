package main

// https://blog.golang.org/comcurrency-timeouts

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	// range
	for data := range c {
		fmt.Println(data)
	}

	// 带缓冲的通道
	cache := make(chan int, 3)

	// 内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(cache), cap(cache))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			cache <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(cache), cap(cache))
		}
	}()

	time.Sleep(2 * time.Second) // 延时2s
	for i := 0; i < 3; i++ {
		num := <-cache // 从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}
	fmt.Println("main进程结束")
}
