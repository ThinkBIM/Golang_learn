package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "www.thinkbim.cn"

	if strings.HasPrefix(myString, "www1") {
		fmt.Println(1)
	}else {
		fmt.Println(0)
	}
}
