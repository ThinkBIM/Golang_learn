package main

import "fmt"

func main()  {
	fc()
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}
func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func fc() {
	for i := 0; i < 5; i++ {
		//关键字 defer 允许我们进行一些函数执行完成后的收尾工作
		defer fmt.Printf("%d ", i)
	}
}