package main

import "fmt"

func main() {

	fmt.Println("General（通用占位符）------")
	fmt.Printf("%%v 以默认的方式打印变量的值   %v\n", 1.2)
	fmt.Printf("%%T 打印变量的类型   %T\n", 1.2)
	fmt.Printf("%% 字面上的百分号，并非值的占位符\n")


	fmt.Println("------Integer（整型）------")
	fmt.Printf("%%+d 带符号的整型 %+d\n", -11)
	fmt.Printf("%%q 打印单引号 %q\n", "'222'")
	fmt.Printf("%%o 不带零的八进制 %o\n", 100)
	fmt.Printf("%%#o 带零的八进制 %#o\n", 100)
	fmt.Printf("%%x 小写的十六进制 %x\n", 100)
	fmt.Printf("%%X 大写的十六进制 %X\n", 100)
	fmt.Printf("%%#x 带0x的十六进制 %#x\n", 100)
	fmt.Printf("%%U 打印Unicode字符 %U\n", 100)
	fmt.Printf("%%#U 打印带字符的Unicode %#U\n", 100)
	fmt.Printf("%%b 打印整型的二进制 %b\n", 100)


	fmt.Println("------Integer width（指定长度的整型，以5为例）------")
	fmt.Printf("%%5d 整型长度为5，右对齐，左边留白 %5d\n", 1000)
	fmt.Printf("%%-5d 左对齐右边留白 %-5d\n", 1000)
	fmt.Printf("%%05d 数字前面补零 %05d\n", 1000)


	fmt.Println("------Float（浮点数）------")
	fmt.Printf("%%f 6位小数点(=%%.6f)   %f\n", 1.2)
	fmt.Printf("%%e 6位小数点（科学计数法）(=%%.6e)   %e\n", 1.2)
	fmt.Printf("%%g 用最少的数字来表示   %g\n", 1.2000)
	fmt.Printf("%%.3g 最多3位数字来表示   %.3g\n", 1.2000)
	fmt.Printf("%%.3f 最多3位小数来表示   %.3f\n", 1.2000)

	fmt.Println("------String Width (指定长度的字符串）------")
	str := "人者多欲，其性尚私。成事享其功，败事委其过。\n\n人心多诈,不可视其表。信人莫若信己,防人毋存幸念。"
	fmt.Printf("%%5s 打印字符串，最小宽度为5  %5s\n", "123")
	fmt.Printf("%%-5s 打印字符串，最小宽度为5（左对齐）  %-5s\n", "123")
	fmt.Printf("%%.5s 打印字符串，最大宽度为5  %.5s\n", "123")
	fmt.Printf("%%5.7s 打印字符串，最小宽度为5，最大宽度为7 %5.7s\n", str)
	fmt.Printf("%%-5.7s 打印字符串，最小宽度为5，最大宽度为7（左对齐） %-5.7s\n", str)
	fmt.Printf("%%5.3s 打印字符串，如果宽度大于3，则截断 %5.3s\n", str)
	fmt.Printf("%%05s 打印字符串，如果宽度小于5，就会在字符串前面补零 %05s\n", "123")



	fmt.Println("------Struct（结构体）------")
	i := "北京"
	fmt.Printf("%%v 以默认的方式打印变量的值 %v\n", i)
	fmt.Printf("%%+v 在打印结构体时，会添加字段名 %+v\n", i)
	fmt.Printf("%%#v 在打印结构体时，会添加字段名和包名 %#v\n", i)


	fmt.Println("------Boolean（布尔值）------")
	fmt.Printf("%%t 打印布尔值  %t\n", true)

	fmt.Println("------Pointer（指针）------")
	fmt.Printf("%%p 打印指针  %p\n", i)
	fmt.Printf("%%#p 不带0x的指针  %#p\n", i)



	fmt.Printf("%%U 打印带字符的Unicode  %s\n", "\u4e2d\u6587")
}
