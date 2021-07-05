package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "你好，Go"
	fmt.Printf("string : %v\n", str)

	input := []byte(str)
	fmt.Printf("byte : %v\n", input)

	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("encode base64 : %v\n", encodeString)

	decodeBytes, _ := base64.StdEncoding.DecodeString(encodeString)
	fmt.Printf("decode base64 : %v\n", string(decodeBytes))

	urlStr := "?word=go&text=你好呀这是怎么了"
	urlencode := base64.URLEncoding.EncodeToString([]byte(urlStr))
	fmt.Printf("urlencode : %v\n", urlencode)

	urldecode, _ := base64.URLEncoding.DecodeString(urlencode)
	fmt.Printf("urldecode : %v\n", string(urldecode))
}
