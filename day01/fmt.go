package main

import "fmt"

//go语言推荐用驼峰命名法
//非全局变量声明后必须使用

var hello string

func main() {
	// fmt.Print("A")
	// fmt.Print("B")
	// fmt.Print("C")

	// fmt.Println("A")
	// fmt.Println("B")
	// fmt.Println("C")

	// var user = "张三"
	// fmt.Println(user)
	// var a int = 123
	// user = "李四"
	// fmt.Println(user)
	// fmt.Printf("%v", a)

	var (
		name string
		age  int
		isOk bool
	)
	name = "张三"
	age = 20
	isOk = true

	fmt.Println(name, age, isOk)
}
