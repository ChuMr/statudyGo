package main

import "fmt"

//空接口的应用
//类型断言
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {

	//map 的value值 ，函数不知道传什么参数时
	// var m1 map[string]interface{}
	// m1 = make(map[string]interface{}, 10)

	m1 := make(map[string]interface{}, 10)
	m1["name"] = "zhangsan"
	m1["age"] = 1000
	m1["hobby"] = [...]string{"唱歌", "跳舞", "rap"}
	m1["id"] = [...]int{12, 132, 342}
	m1["t"] = true
	show(true)
	show(m1["name"])
	show(m1)

}
