package main

import "fmt"

func main()  {

	for i:=1;i<=15;i++ {
		fmt.Println(i)
	}
	

	//var a string = "G"
	//for i:=1;i<=25;i++ {
	//	a += a
	//	fmt.Printf(a)
	//}


	for s := ""; s != "aaaaaaaaaa"; {
		fmt.Println(s)
		s = s +"a"
	}

	for i, j, s :=0,5,"a"; i <3&& j <100&& s !="aaaaa"; i, j,
		s = i+1, j+1, s +"a"{
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	str2 := "你好呀"

	for index, rune := range str2 {
		fmt.Printf("%c", rune)
		fmt.Println()
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune,[]byte(string(rune)))
	}


	for i:=0; i<4; i++ {
		for j:=0; j<10; j++ {
			if j>5 {
				break
			}
			print(j)
		}
		print("  ")
	}

	fmt.Println()

	c := 0
	HEAR:
		print(c)
		c++
		if c == 5 {
			return
		}
		goto HEAR
}