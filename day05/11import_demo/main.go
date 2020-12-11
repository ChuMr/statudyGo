package main

import (
	"fmt"
	// calc "../10calc"
	calc "zhangqian.com/day05/10calc"
)

func init() {
	fmt.Println("import_demo 初始化")
	fmt.Println(name)
}

var name = "张三"

func main() {
	// calc "zhangqian.com/day05/10calc"
	fmt.Println("start")
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
