package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"os"
)

func main() {
	yml := "/Users/zhangcheng/WWW/Golang_learn/package/yaml/data.yml"
	r, _ := os.ReadFile(yml)
	js, _ := yaml.YAMLToJSON(r)
	// j := []byte(`{"name": "John", "age": 30}`)
	// y, _ := yaml.JSONToYAML(j)
	fmt.Println(string(js))
	// ApplyYAML(yml)
}

// func ApplyYAML(yml string) {
// 	r, err := os.Open(yml)
// 	if err != nil {
// 		fmt.Println("file open err", err)
// 	}
// 	defer r.Close()
//
// 	dec := yaml.NewDecoder(r)
//
// 	fmt.Println(dec)
// }
