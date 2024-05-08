package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	var u User
	u.Age = 10
	u.Name = "sssss"
	//h := `{"name":"张三","age":15}`
	//err := json.Unmarshal([]byte(h), &u)
	//if(err != nil) {
	//	fmt.Println(err)
	//}
	fmt.Printf(u.Name)
}

func test()  {
	var u User
	t := reflect.TypeOf(u)

	for i:=0; i<t.NumField();i++ {
		sf:=t.Field(i)
		//fmt.Println(sf)
		fmt.Println(sf.Tag)
	}
}