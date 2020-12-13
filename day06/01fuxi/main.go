package main

import "fmt"

func main() {
	var a interface{}
	// var a int32
	// var a int8
	// a = 100
	// a = int32(100)
	a = int64(100)
	// a := 100
	// fmt.Println(a)
	// v, ok := a.(int8)
	// fmt.Println(v)
	// fmt.Println(ok)

	switch v1 := a.(type) {
	case int:
		fmt.Printf("类型int,值为:%v", v1)
	case int32:
		fmt.Printf("类型int32,值为:%v", v1)
	default:
		fmt.Println("其他类型")
	}
}
