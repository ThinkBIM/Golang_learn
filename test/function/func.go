package main

import "fmt"

var num int=10
var numx2, numx3 int

func main(){
	b()
	//function1()
	//fmt.Println、(match(1, 2))
	//ln(1,2,3,4,5)
	//numx2, numx3 = getX2AndX3(num)
	//PrintValues()
	//numx2, numx3 = getX2AndX3_2(num)
	//PrintValues()
	//PrintValues()
}
func PrintValues(){
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
func getX2AndX3(input int)(int,int){
	return 2 * input,3 * input
}
func getX2AndX3_2(input int)(x2 int, x3 int){
	x2 = 4 * input
	x3 = 6 * input
	// return x2, x3
	return
}

func match(a int, b int) (int, int, int) {
	return a+b, a*b, a-b
}

func ln(arg ...int)  {
	for _, value := range arg {
		fmt.Println(value)
	}
	fmt.Println(arg)
}




func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }
func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}
func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}


//func (st *Stack) Pop() int {
//	v := 0
//	for ix := len(st) - 1; ix >= 0; ix-- {
//		if v = st[ix]; v != 0 {
//			st[ix] = 0
//			return v
//		}
//	}
//	return v
//}