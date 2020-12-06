package main

import "fmt"

//给自定义类型加方法

type myInt int

func (m myInt) hello() {
	fmt.Println("test")
}
func main() {
	m := myInt(100)
	m.hello()
	fmt.Println(m)
}
