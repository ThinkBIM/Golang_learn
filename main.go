package main

import (
	"Golang_learn/funcs"
	"fmt"
)

func main() {

	m := make(map[string]map[string]interface{})

	m["s"] = map[string]interface{}{
		"apple":  5,
		"banana": "sdfsd",
		"cherry": "fs",
	}

	m["b"] = map[string]interface{}{
		"apple":  15,
		"banana": "12",
		"cherry": "fss",
	}
	res := funcs.ArrayColumn(m, "apple")

	fmt.Println(res)
}
