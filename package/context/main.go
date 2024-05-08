package main

import (
	"context"
	"fmt"
	"time"
)

// WithTimeout：一般用户控制超时
// WithCancel：用于取消整条链上的任务

// Backgroud：返回一个空的 context.Context
// ToDo：返回一个空的 context.Context，但是
// 这个标记着你也不知道传什么

// WithDeadline 控制时间
func WithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()

	fmt.Printf("%v-%v", start, end)
}

// WithValue 往里面塞入 key-value
func WithValue() {
	parentKey := "parent"
	parent := context.WithValue(context.Background(), parentKey, "this is parent")

	sonKey := "son"
	son := context.WithValue(parent, sonKey, "this is son")

	fmt.Printf("%v \n", parent.Value(parentKey))
	fmt.Printf("%v \n", parent.Value(sonKey)) // nil 访问不到son

	fmt.Printf("%v \n", son.Value(parentKey))
	fmt.Printf("%v \n", son.Value(sonKey))
}

func main() {
	WithDeadline()
	WithValue()
}
