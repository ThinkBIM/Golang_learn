package main

import (
	"fmt"
	"log"
	"runtime"
)

func main()  {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	
	where()

	for i := 0; i < 4; i++ {
		tests(i)
	}

	log.SetFlags(log.Llongfile)
	log.Print("111")
}

func tests(skip int)  {
	call(skip)
}

func call(skip int)  {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name()
	fmt.Println(fmt.Sprintf("%v %s %d %t %s", pc, file, line, ok, pcName))
}