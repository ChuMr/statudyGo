package main

import (
	"fmt"
	"unicode"
)

func main() {
	//判断字符串中，汉子的数量
	//
	s := "hello 上海shd是"
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
	//how do you do 单词出现的次数
}
