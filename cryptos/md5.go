package cryptos

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(str string) {
	sum := md5.Sum([]byte(str))
	fmt.Printf("md5 : %v\n", hex.EncodeToString(sum[:]))
}
