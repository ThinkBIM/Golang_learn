package goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// type Comic struct {
// 	Id        int    `json:"id"`
// 	Name      string `json:"name"`
// 	Cover     string `json:"cover"`
// 	DetailUrl string `json:"detail_url"`
// }

type Contents struct {
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

// func (c *Comic) SetCover(str string) {
// 	c.Cover = str
// }

// func top(page int) {
// 	res, err := http.Get("https://www.agemys.com/rank?tag=catyear&value=&page=" + strconv.Itoa(page))
// 	if err != nil {
//
// 	}
// 	defer res.Body.Close()
//
// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	if err != nil {
//
// 	}
//
// 	var comics []Comic
//
// 	doc.Find(".div_right_r").Each(func(i int, s *goquery.Selection) {
//
// 		// fmt.Printf("div_right_r : %d\n", i)
// 		s.Find(".rank_text").Each(func(ii int, ss *goquery.Selection) {
// 			// fmt.Printf("rank_text: %d\n", ii)
// 			content := ss.Find(".rank_text_name").Text()
// 			id := ss.Find(".rank_text_idx").Text()
// 			attr, _ := ss.Find("a").Attr("href")
//
// 			idd, _ := strconv.Atoi(id)
//
// 			comics = append(comics, Comic{
// 				Id:        idd,
// 				Name:      content,
// 				Cover:     "",
// 				DetailUrl: attr,
// 			})
//
// 			// fmt.Printf("%s: %s  %s\n", id, content, attr)
// 		})
// 	})
//
// 	for k, i2 := range comics {
// 		v, err := http.Get("https://www.agemys.com" + i2.DetailUrl)
// 		if err != nil {
// 			return
// 		}
// 		doc, err := goquery.NewDocumentFromReader(v.Body)
// 		if err != nil {
//
// 		}
// 		attr, _ := doc.Find(".baseblock").Eq(0).Find("img").Attr("src")
// 		if attr != "" {
// 			if strings.HasPrefix(attr, "//") {
// 				attr = "https:" + attr
// 			}
// 			comics[k].Cover = attr
// 		}
//
// 		v.Body.Close()
//
// 		// fmt.Printf("%v", i2)
// 		// return
// 	}
// 	// data, _ := json.Marshal(comics)
// 	fmt.Printf("%v", comics)
// }

// NewsSinaList 抓取新浪新闻数据
func NewsSinaList() {
	url := "https://tzxy.sina.com.cn/kuaixun"
	res, err := http.Get(url)
	if err != nil {

	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {

	}
	fmt.Println(doc)
	fmt.Println(11111111111)

	// var conts []Contents

	doc.Find("#liveList01").Each(func(i int, selection *goquery.Selection) {
		// name := selection.Find(".anime_info h3").Text()

		// fmt.Println(11111)
		fmt.Println(i)
		fmt.Println(selection.Html())
		// return
		// cover, _ := selection.Find(".overflow-hidden.anime_cover_image img").Attr("src")
		// conts = append(conts, Contents{
		// 	Name: name,
		// })
	})

	// str, _ := json.Marshal(newComic)
	//
	// fmt.Printf("%s", str)
	// return

	// for i, comic := range newComic {
	// 	filename := "/Users/zhangcheng/WWW/漫画/202301/" + fmt.Sprintf("%d.jpg", i)
	// 	err := kpl.DownImg(filename, comic.Cover)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}
