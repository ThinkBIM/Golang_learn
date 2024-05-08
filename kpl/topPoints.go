package kpl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func js() {
	v, _ := http.Get("https://pvp.qq.com/match/kpl/kingproleague/js/config-data.js")

	// var buf = make([]byte, 4096)
	// var data []byte

	defer v.Body.Close()

	// n, err := v.Body.Read(data)
	// if err != nil {
	//	return
	// }
	// if n == 0 {
	//	return
	// }
	// for {
	//
	//	//data = append(data, buf[:n]...)
	// }
	// decoder := m

	body, _ := ioutil.ReadAll(v.Body)

	fmt.Println(string(body))
}

func main() {

	js()

	// v, _ := http.Get("https://pvp.qq.com/match/kpl/kingproleague/index.html")
	//
	// if v.StatusCode != 200 {
	//	fmt.Println("请求失败")
	// }
	//
	// doc, _ := goquery.NewDocumentFromReader(v.Body)
	//
	// html, _ := doc.Find("#sszb").Html()
	//
	// fmt.Println(html)
}
