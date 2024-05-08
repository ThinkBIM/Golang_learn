package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLine(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func main() {
	r, _ := os.Open("/Users/zhangcheng/WWW/doc_help/1.txt")
	// fmt.Printf("%v", r)
	line, err := CountLines(r)
	fmt.Println(line)
	fmt.Println(err)
}
