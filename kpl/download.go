package kpl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func download(heroId int, index int, pic string) error {
	url := fmt.Sprintf("http://game.gtimg.cn/images/yxzj/img201606/skin/hero-info/%d/%d-bigskin-%d.jpg", heroId, heroId, index)
	fmt.Println(url)

	filename := pic + fmt.Sprintf("%d_%d.jpg", heroId, index)

	// 检查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		v, err := http.Get(url)
		if err != nil {
			fmt.Println("http err", err)
			return err
		}
		defer v.Body.Close()

		content, err := ioutil.ReadAll(v.Body)
		if err != nil {
			fmt.Println("err", err)
			return err
		}

		err = ioutil.WriteFile(filename, content, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

func DownImg(filename string, url string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		v, err := http.Get(url)
		if err != nil {
			fmt.Println("http err", err)
			return err
		}
		defer v.Body.Close()

		content, err := ioutil.ReadAll(v.Body)
		if err != nil {
			fmt.Println("read err", err)
			return err
		}
		err = ioutil.WriteFile(filename, content, 0666)
		if err != nil {
			fmt.Println("write err", err)
			return err
		}
	}
	return nil
}

// func main() {
//	//s()
//	//i := 1
//	//heroId := 542
//	//pic := "/Users/zhangcheng/WWW/kpl/skin/"
//	//
//	//err := download(heroId, i, pic)
//	//if err != nil {
//	//	return
//	//}
// }
