package main

import (
	"fmt"
	"time"
)

func channelWithCache() {
	// 缓存
	ch := make(chan string, 1)
	go func() {
		ch <- "第一次"
		time.Sleep(time.Second)
		ch <- "第二次"
	}()

	time.Sleep(2 * time.Second)
	msg := <-ch
	fmt.Println(time.Now().String() + msg)

	msg = <-ch

	fmt.Println(time.Now().String(), msg)

}

func channelWithoutCache() {
	// 不缓存
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "第一次"
	}()

	msg := <-ch
	fmt.Println(msg)
}

func main() {
	channelWithCache()
	channelWithoutCache()
}
