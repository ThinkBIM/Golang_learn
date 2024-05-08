package main

import "fmt"

/**
// 变量声明  全局变量 常量
//
//
*/

// A 声明全局变量 := 不支持全局变量声明，只能在方法中使用
var A = 100
var B int = 100

//常量, 不允许修改
const length int = 10

const (
	BEIJING  = iota //关键字 iota,只能在 const中使用 每行初始 0，没一行的iota都会加1
	SHANGHAI        //iota 1
	TIANJIN
	TIANJI
	TIANJ
	TIAN
	TIA
	TI
	T
)

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := 3.8
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	//声明多变量
	var x, y int = 100, 200
	var xx, yy = 1, "22a"
	fmt.Printf("x = %v\n y = %v\n", x, y)
	fmt.Printf("xx = %v\n yy = %v\n", xx, yy)

	var (
		aa     = 100
		bb     = "nis"
		cc int = 200
	)
	fmt.Printf("aa = %v\n bb = %v\n cc = %v\n", aa, bb, cc)

	fmt.Println("BEIJING", BEIJING)
	fmt.Println("SHANGHAI", SHANGHAI)
	fmt.Println("TIANJIN", TIANJIN)
	fmt.Println("TIANJI", TIANJI)
	fmt.Println("TIANJ", TIANJ)
	fmt.Println("TIAN", TIAN)
	fmt.Println("TIA", TIA)
	fmt.Println("TI", TI)
	fmt.Println("T", T)
}
