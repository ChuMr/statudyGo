package main

import "fmt"

//painc recover
func main() {
	// 获取用户输入
	var s string
	fmt.Scan(&s)

	fmt.Println(s)
	// fA()
	// fB()
	// fC()
}

func fA() {
	fmt.Println("a")
}

func fB() {
	fmt.Println("b")
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("error")
}

func fC() {
	fmt.Println("c")
}
