package main

import (
	"fmt"
)

func main() {
	//多行的字符串

	// var b1 = "hello word"
	// var b2 = 'h'
	// var b3 = `测hi是是  是
	// 是
	// 是
	// 是`
	// fmt.Println(b1)
	// fmt.Println(b2)
	// fmt.Println(b3)

	// //字符串拼接
	// var c = "hello"
	// c1 := "word"
	// c2 := c + c1
	// fmt.Println(c2)

	// //分割
	// p := "D:\\GoWork\\src\\qian.zhang.com\\studygo"
	// ret := strings.Split(p, "\\")
	// fmt.Println(ret)

	// e := strings.Contains(p, "Go1")
	// fmt.Println(e)

	// //前缀
	// fmt.Println(strings.HasPrefix(p, "D:"))

	// fmt.Println(strings.Index(p, "s"))
	// fmt.Println(strings.LastIndex(p, "s"))

	// //拼接
	// fmt.Println(strings.Join(ret, ","))

	// h := "hello word 是收到"
	// for k, v := range h {
	// 	fmt.Printf("%d %c\n", k, v)
	// }
	// i := 1
	// j := 1

	// // s := i + j
	// // s1 := fmt.Sprintf("%d*%d=%d \n", i, j, i*j)
	// s1 := fmt.Sprintf("%d*%d=%d ", i, j, i*j)
	// s1 := fmt.Sprintf("%d*%d=%d ", i, j, i*j)
	// // s := i + '*' + j + '=' + i*j + '-'
	// fmt.Print(s1)

	//99 乘法表
	for i := 1; i <= 9; i++ {
		var s string
		for j := 1; j <= i; j++ {
			s = fmt.Sprintf("%d*%d=%d ", i, j, i*j)
			fmt.Print(s)
		}
		fmt.Println()
	}
}
