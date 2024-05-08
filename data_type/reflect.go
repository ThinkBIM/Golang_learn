// reflect包的应用
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

// 结构体标签
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

// Movie 结构体标签 JOSN
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float32  `json:"price"`
	Actors []string `json:"actors"`
}

func (u User) Say() {
	fmt.Println("User method Say ...")
	fmt.Printf("%v\n", u)
}

func (u User) GetName() {
	fmt.Println("User method GetName ...")
	fmt.Printf("%v\n", u)
}

func main() {
	a := 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

	user := User{"zhang3", 18, "男"}
	DoFeachFieldMetchod(user)

	var r resume
	FindTag(&r)

	// json
	movie := Movie{
		"喜剧之王",
		2000,
		68.88,
		[]string{"周星驰", "张柏芝"},
	}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("struct to json error", err)
		return
	}
	fmt.Printf("json %s\n", jsonStr)

	// struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json to struct error", err)
		return
	}
	fmt.Printf("struct %v\n", myMovie)
}

func DoFeachFieldMetchod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	// 获取所有所有字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s  %v  %v\n", field.Name, field.Type, value)
	}
	// 获取所有方法
	for i := 0; i < inputType.NumMethod(); i++ {
		metchod := inputType.Method(i)
		fmt.Printf("%s  %v\n", metchod.Name, metchod.Type)
	}
}

// FindTag 获取结构体的所有标签
func FindTag(str interface{}) {
	// 获取所有便签
	tag := reflect.TypeOf(str).Elem()
	for i := 0; i < tag.NumField(); i++ {
		tagInfo := tag.Field(i).Tag.Get("info")
		tagDoc := tag.Field(i).Tag.Get("doc")
		fmt.Println("Tag info ", tagInfo, "Tag doc ", tagDoc)
	}
}
