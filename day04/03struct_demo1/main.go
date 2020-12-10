package main

import "fmt"

//结构体

type preson struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p preson
	p.name = "这是个人"
	p.age = 12
	p.gender = "男"
	p.hobby = []string{"C++", "PHP", "golang", "Python"}
	// fmt.Println(p)

	//匿名结构体
	var s struct {
		name string
		age  int
	}
	s.name = "test"
	s.age = 1000
	fmt.Println(s)
	fmt.Printf("type:%T value:%v\n", s, s)
}
