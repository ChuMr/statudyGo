package main

import "fmt"

func main() {
	ss := "上海自来水来自海上"
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
}
 