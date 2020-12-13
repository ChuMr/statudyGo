package main

import (
	"fmt"
)

//类型断言

func assign(a interface{}) {
	str, ok := a.(int64)
	fmt.Println(str)
	fmt.Println(ok)
}

// func assign2(a interface{}) {
// 	switch t := a.(type) {
// 	case string:
// 		fmt.Println("字符串:", t)
// 	case int:
// 		fmt.Println("int:", t)
// 	case int32:
// 		fmt.Println("int32:", t)
// 	case int64:
// 		fmt.Println("int64:", t)
// 	case bool:
// 		fmt.Println("bool:", t)
// 	default:
// 		fmt.Println("bool:", t)
// 	}
// }

func main() {
	var a = int64(100)
	fmt.Printf("类型：%T 值：%v", a, a)
	// assign(a)
	// a := int(100)
	// str, ok := 100.(int)
	// fmt.Println(str)
	// fmt.Println(ok)
	// t := a.(type)
	// switch t := a.(type) {
	// case int:
	// 	fmt.Println("int~~~~~~~~~")
	// }
	// fmt.Println(t)
	// fmt.Println(ok)
	// assign(100)
	// i := int64(200)
	// assign2(100)
	// assign2(int64(100))
	// assign2(false)
}
