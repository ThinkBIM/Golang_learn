package main

import "fmt"

func foo1(a int, b string) int {
	return 1
}

// 返回多参数，匿名
func foo2(a int, b string) (int, int) {
	return 1, 2
}

// 返回多参数，有形参名称
func foo3(a int, b int) (a1 int, b1 int) {
	// 形参赋值
	a1 = 100
	b1 = 200
	return
}

func foo4(a int, b int) (a1, b1 int) {

	return
}

func main() {
	var a *string
	cc := "10"
	a = &cc
	fmt.Printf("%v", *a)

}
