package main

import "fmt"

type preson struct {
	name string
	age  int
}

//构造函数
func newPreson(name string, age int) *preson {
	return &preson{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPreson("test1", 100)
	fmt.Println(p1)
}
