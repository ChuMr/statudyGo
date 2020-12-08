package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

type chicken struct {
	feet int
}

func (c *cat) move() {
	fmt.Println("猫喵喵喵")
}

func (c *cat) eat(foot string) {
	fmt.Printf("猫吃%s\n", foot)
}

func (c chicken) move() {
	fmt.Println("鸡咕咕咕")
}

func (c chicken) eat(foot string) {
	fmt.Printf("鸡吃%s", foot)
}

func main() {

	var a animal

	cat1 := cat{
		"蓝猫",
		4,
	}

	// cat2 := cat{
	// 	"红猫",
	// 	4,
	// }

	a = &cat1
	fmt.Println(a)
	fmt.Println(cat1)
	// fmt.Println(cat2)

	// fmt.Printf("值%v,类型%T\n", a, a)

	// c1 := cat{
	// 	name: "蓝猫",
	// 	feet: 4,
	// }
	// fmt.Printf("值%v,类型%T\n", c1, c1)

	// a = c1
	// fmt.Printf("值%v,类型%T\n", a, a)

	// c2 := chicken{
	// 	feet: 4,
	// }
	// c2.eat("虫子")
	// a = c2
	// fmt.Printf("a值%v,a类型%T\n", a, a)
}
