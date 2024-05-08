package cryptos

import (
	"fmt"
	"net/url"
)

func Urlc(str string) {
	urlencode := url.QueryEscape(str)
	fmt.Printf("urlencod : %v\n", urlencode)

	urldecode, _ := url.QueryUnescape(urlencode)
	fmt.Printf("urldecodes : %v\n", urldecode)
}
