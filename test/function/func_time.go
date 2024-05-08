package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func longCalculation()  {
	for i := 0; i < 1e1; i++ {
		fmt.Println(i)
	}
}