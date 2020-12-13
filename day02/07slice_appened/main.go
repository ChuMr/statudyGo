package main

import "fmt"

//append 为切片添加元素
func main() {

	// var s2 = []string{}
	// s1 := []string{"北京", "上海", "深圳"}

	// s2 = append(s1, "广州")
	// fmt.Println(s1)
	// fmt.Println(s2)

	//append 删除某个元素
	a1 := [...]int{1, 3, 6, 2, 8, 0, 3, 12, 9, 54, 32432}
	s1 := a1[:] //数组转成切片

	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1) //   6 2 8 0 3 12 9 54 32432 追加到1的后面，但是底层数据 索引10 为32432的数据还在，结果就是[1 6 2 8 0 3 12 9 54 32432 32432]

}
