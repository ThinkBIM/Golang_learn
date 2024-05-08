package main

import "fmt"

//匿名函数
//https://www.bookstack.cn/read/the-way-to-go_ZH_CN/eBook-06.9.md
func main()  {
	fplus := func(x, y int) int { return x + y }
	fmt.Println(fplus(1,2))

	//表示参数列表的第一对括号必须紧挨着关键字 func，因为匿名函数没有名称。花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。
	func() {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()
	fmt.Println()

	f()
	fmt.Println(c())
}

func f()  {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

func c() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

//闭包函数：将函数作为返回值



