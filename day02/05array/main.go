package main

import "fmt"

//数组 数组的长度是数组类型的一部分
func main() {
	// var a1 [3]string
	// var a2 [3]bool
	// // var a3 [3]int

	// //1数组的初始化
	// a1 = [3]string{"a", "b", "c"}

	// //2根据初始值自动推算数组的长度
	// a3 := [...]int{1, 2}

	// //初始化3根据索引来初始化
	// a5 := [5]int{3, 5, 6}
	// a6 := [5]int{0: 1, 4: 2}
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println(a5)
	// fmt.Println(a6)
	// //多维数组
	// a10 := [3][2]string{
	// 	{"a1", "a2"},
	// 	{"a1", "a2"},
	// 	{"a1", "a2"},
	// }
	// fmt.Println(a10)
	//数组的遍历
	//数组是值类型不是引用类型

	//对数组所有元素求和 [1,432,432,1,23,75,8]
	a20 := [...]int{1, 54, 21, 1, 23, 75, 8}
	// var total int
	// for i := 0; i < len(a20); i++ {
	// 	total += a20[i]
	// }
	// fmt.Println(total)

	//找出数组指定值的两个元素的下标
	var flag = false
	for i := 0; i < len(a20); i++ {
		for j := i + 1; j < len(a20); j++ {
			if a20[i]+a20[j] == 9 {
				fmt.Println(a20[i], a20[j])
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
}
