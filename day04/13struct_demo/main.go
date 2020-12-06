package main

import "fmt"

/*
 * 匿名字段 and 嵌套结构体
 */
//嵌套结构体
type address struct {
	provice string
	city    string
}
type preson struct {
	name string
	age  int
	addr address
}

type company struct {
	name string
	address
}

func main() {

	// p1 := preson{
	// 	"张三",
	// 	18,
	// }
	// fmt.Println(p1)
	// fmt.Println(p1.string) //匿名字段
	p2 := preson{
		name: "张三",
		age:  18,
		addr: address{
			provice: "上海",
			city:    "静安",
		},
	}
	fmt.Println(p2)
	fmt.Println(p2.name)
	fmt.Println(p2.addr.city)

	c1 := company{
		name: "国际公司",
		address: address{
			provice: "香港",
			city:    "尖沙咀",
		},
	}
	fmt.Println(c1)
	fmt.Println(c1.name)
	fmt.Println(c1.city)
}
