package main

import "fmt"

//vscode 不支持gomodule
//指针
func main() {
	// & 取地址
	// * 根据地址取值
	// n := 18
	// p := &n //内存地址
	// fmt.Printf("%T\n", p)

	// m := *p
	// fmt.Println(m)

	var a1 *int
	var a2 = new(int)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(a1)
	fmt.Println(a2)

}
