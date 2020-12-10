package main

import (
	"fmt"
)

//第一天复习
var age = 18
var name string

func main() {
	// var isOk = true
	// fmt.Println(isOk)
	// fmt.Println(age)
	// fmt.Println(name)
	//变量和常量

	// var name1 string
	// var name2 = "go"
	// name3 := "python" //只能在函数内部使用

	//匿名变量 _
	//常量const Pi = 3.1415926
	//iota const 关键字出现的时候被重置为0， const 每新增一行常量声明，iota +1，注意是每新增一行

	//流程控制 和 for
	// if f := 1; f < 10 {
	// 	fmt.Println("小于10")
	// } else {
	// 	fmt.Println("大于10")
	// }

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}

	//倒得
	// for i := 9; i >= 1; i-- {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d*%d=%d\t", j, i, i*j)
	// 	}
	// 	fmt.Println()
	// }
	//整型unit 无符号  int有符号
	//

	// 浮点型 float64  float32
	// 复数 complex128 complex64
	// 布尔值  true false 不能和其他类型做运算
	// 字符串
	// byte rune 类型
	// 字符串 golang 用双引号； 字符用单引号； 字节 1 byte = 8 bit

	// byte rune(int32)
	// s1 = "hello"
	// s2 = "上海有上又有海"

	for i := 0; i < 10; i++ {
		if i == 2 {
			goto le
		}
	}
	fmt.Println("结束")
le:
	fmt.Println("le")
}
