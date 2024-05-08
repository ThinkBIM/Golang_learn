package goquery

import (
	"Golang_learn/kpl"
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Cover     string `json:"cover"`
	DetailUrl string `json:"detail_url"`
}

type NewComic struct {
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

func (c *Comic) SetCover(str string) {
	c.Cover = str
}

func top(page int) {
	res, err := http.Get("https://www.agemys.com/rank?tag=catyear&value=&page=" + strconv.Itoa(page))
	if err != nil {

	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {

	}

	var comics []Comic

	doc.Find(".div_right_r").Each(func(i int, s *goquery.Selection) {

		// fmt.Printf("div_right_r : %d\n", i)
		s.Find(".rank_text").Each(func(ii int, ss *goquery.Selection) {
			// fmt.Printf("rank_text: %d\n", ii)
			content := ss.Find(".rank_text_name").Text()
			id := ss.Find(".rank_text_idx").Text()
			attr, _ := ss.Find("a").Attr("href")

			idd, _ := strconv.Atoi(id)

			comics = append(comics, Comic{
				Id:        idd,
				Name:      content,
				Cover:     "",
				DetailUrl: attr,
			})

			// fmt.Printf("%s: %s  %s\n", id, content, attr)
		})
	})

	for k, i2 := range comics {
		v, err := http.Get("https://www.agemys.com" + i2.DetailUrl)
		if err != nil {
			return
		}
		doc, err := goquery.NewDocumentFromReader(v.Body)
		if err != nil {

		}
		attr, _ := doc.Find(".baseblock").Eq(0).Find("img").Attr("src")
		if attr != "" {
			if strings.HasPrefix(attr, "//") {
				attr = "https:" + attr
			}
			comics[k].Cover = attr
		}

		v.Body.Close()

		// fmt.Printf("%v", i2)
		// return
	}
	// data, _ := json.Marshal(comics)
	fmt.Printf("%v", comics)
}

// Acgsecrets 抓取新番动漫
// https://acgsecrets.hk/bangumi/202207/
func Acgsecrets() {
	url := "https://acgsecrets.hk/bangumi/202404/"
	res, err := http.Get(url)
	if err != nil {

	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {

	}

	var newComic []NewComic

	doc.Find("div[acgs-bangumi-anime-id]").Each(func(i int, selection *goquery.Selection) {
		name := selection.Find(".anime_info h3").Text()
		// fmt.Println(name)
		// return
		cover, _ := selection.Find(".overflow-hidden.anime_cover_image img").Attr("src")
		newComic = append(newComic, NewComic{
			Name:  name,
			Cover: cover,
		})
	})

	filePath := "/Users/zhangcheng/Desktop/202401/title.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 及时关闭file句柄
	defer file.Close()
	// 写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	for i, comic := range newComic {
		filename := "/Users/zhangcheng/Desktop/202401/" + fmt.Sprintf("%d.jpg", i)
		err := kpl.DownImg(filename, comic.Cover)
		if err != nil {
			fmt.Println(err)
		}
		write.WriteString(comic.Name + "\r\n")
	}
	// Flush将缓存的文件真正写入到文件中
	write.Flush()
}
