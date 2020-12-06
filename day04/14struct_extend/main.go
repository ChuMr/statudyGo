package main

import "fmt"

//模拟面向对象继承
//利用嵌套模拟面向对象继承

type animal struct {
	name string
}

type dog struct {
	feet int
	animal
}

func (d dog) wang() {
	fmt.Printf("%s再叫，汪汪汪", d.name)
}

func (a animal) move() {
	fmt.Printf("%s 会动", a.name)
}

func main() {
	d1 := dog{
		feet: 3,
		animal: animal{
			name: "小狗",
		},
	}

	d1.wang()
}
