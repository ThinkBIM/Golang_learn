package goquery

import (
	"Golang_learn/kpl"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

type NewGuiqi struct {
	Cover string `json:"cover"`
}

func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	// 初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	// 创建一个上下文，超时时间为40s  此时间可做更改  调整等待页面加载时间
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		// log.Fatal("Run err : %v\n", err)
		return "", err
	}
	// log.Println(htmlContent)

	return htmlContent, nil
}

var newGuiqi []NewGuiqi

func GetSpecialData(htmlContent string, selector string) ([]NewGuiqi, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return newGuiqi, err
	}
	var str string
	dom.Find(".comicpage img").Each(func(i int, selection *goquery.Selection) {
		str, _ = selection.Attr("data-original")
		newGuiqi = append(newGuiqi, NewGuiqi{
			Cover: str,
		})
	})
	return newGuiqi, nil
}

// Acgsecrets 抓取新番动漫
// https://acgsecrets.hk/bangumi/202207/
func Chapter() {

	// https://www.patternrecognition.cn/chapter/567/105903.html
	url := "https://www.patternrecognition.cn/chapter/567/105906.html"
	selector := "#content > div.comiclist"
	param := `document.querySelector("#content > div.comiclist > div")`
	html, _ := GetHttpHtmlContent(url, selector, param)
	res, _ := GetSpecialData(html, "img")

	for i, comic := range res {
		filename := "/Users/zhangcheng/Desktop/地狱鬼妻/1-3/" + fmt.Sprintf("%d.jpg", i)
		err := kpl.DownImg(filename, comic.Cover)
		if err != nil {
			fmt.Println(err)
		}
	}
}
