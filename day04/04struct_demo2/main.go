package main

import "fmt"

//结构体是值类型
type person struct {
	name   string
	gender string
}

//结构体占用一块连续内存
type x struct {
	a int8
	b int8
	c int8
}

func main() {
	// var p1 = new(person)
	// new 基本数据类型，返回的值是对应类型的指针，
	// make slice map channel 返回对应的类型

	// var p person
	// p.name = "名字"
	// p.gender = "男"
	// fmt.Printf("%T\n", p1)

	// fmt.Println(p1)
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
}
