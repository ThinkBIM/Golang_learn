package main

import (
	"fmt"
	str "strings"
)

func main() {
	// myString := "www.thinkbim.cn"
	myString := "1a我"
	//
	// if str.HasPrefix(myString, "thinkbim") {
	//	fmt.Printf("找得")
	// } else {
	//	fmt.Printf("未找到")
	// }
	fmt.Println("Contains 是否包含", str.Contains(myString, "in"))
	fmt.Println("Count 包含次数", str.Count(myString, "w"))
	fmt.Println("HasPrefix 包含前缀", str.HasPrefix(myString, "w"))
	fmt.Println("HasSuffix 包含后缀", str.HasSuffix(myString, "n"))
	fmt.Println("Index 索引位置", str.Index(myString, "n"))
	fmt.Println("Join 拼接", str.Join([]string{"1", "2"}, "n"))
	fmt.Println("Repeat 复制n次", str.Repeat(myString, 2))
	fmt.Println("Replace 查找替换", str.Replace(myString, "w", "s", 1))
	fmt.Println("Split 字符串转数组", str.Split(myString, "."))
	fmt.Println("ToLower 小写", str.ToLower(myString))
	fmt.Println("ToUpper 大写", str.ToUpper(myString))
	fmt.Println("len 字符串长度", len(myString))
	// fmt.Println("Char 字符串转数组", myString[14])
	fmt.Println("Fields 字符串转数组", str.Fields(myString))
	fmt.Println("TrimSpace 字符去除空格", str.TrimSpace(myString))

	// 示例
	str1 := "中文处理"

	fmt.Println(str.Index(str1, "处"))
	fmt.Println("rune 中文处理", []rune("中文处理"))

	path := "/c/a/b/是/2"
	result := str.LastIndex(path, "/")

	fmt.Println(result)

	pre := []byte(path)[0:result]

	re := []rune(string(pre))

	fmt.Println(len(re))

}
