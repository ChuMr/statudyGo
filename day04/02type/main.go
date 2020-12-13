package main

import "fmt"

//自定义类型和类型别名
//type 后面跟的是类型
type myInt int
type mInt = int

func main() {
	var n myInt
	var m mInt
	n = 100
	m = 200
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)
}
