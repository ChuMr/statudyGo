package main

import "fmt"

//递归
//阶乘
func main() {
	fmt.Println(fac(5))
}

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}
