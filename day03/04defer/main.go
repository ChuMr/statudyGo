package main

import "fmt"

func main() {
	//defer
	// fmt.Println("start")
	// defer fmt.Println("defer")
	// fmt.Println("end")
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	//10 1 + 2
	//10,1,2,3

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	//10, 1, 2, 3
	//20, 0, 2, 2
	//2, 0, 2, 2
	//1, 1, 3, 4
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
