package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	OssCli *oss.Client
	e error
)

func init()  {

}

func ossCline()  {
	OssCli, e := oss.New("", "", "")
}

func main()  {

	imagePath := "/Users/zhangcheng/WWW/Golang_learn/image/"
	imageUrl := "http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"

	filename := "1.jpg"



	res, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("")
		return
	}
	defer  res.Body.Close()

	reder := bufio.NewReaderSize(res.Body, 32 * 1024)

	file, err := os.Create(imagePath + filename)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reder)

	fmt.Printf("Total lenth: %d", written)
}