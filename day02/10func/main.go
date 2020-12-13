package main

import "fmt"

// func sum(x, y int) (ret int) {
// 	ret = x + y
// 	return
// }

// func sum2(x, y int) {
// 	fmt.Println(x + y)
// }

// func func3(x, y int) (ret int, r int) {
// 	ret = x + y
// 	r = x - y
// 	return
// }

//统计字符串中每个字符出现的个数
// func countStr() {
// 	str := "hello word wwehdakdysdsywqey"
// 	m = make(map[string]int, 100)
// 	for i := 0; i < len(str); i++ {
// 		f := fmt.Printf("%q", str[i])
// 		m[f]++

// 	}
// 	// m = make(map[string]int, 10)
// 	// for _, s := range str {
// 	// 	// fmt.Println(s)
// 	// 	m[s]++
// 	// 	// fmt.Printf("%q", s)
// 	// 	// fmt.Printf("%T", s)
// 	// 	// fmt.Println()
// 	// }
// 	fmt.Println(m)

// }
func main() {
	// r := sum(20, 10)
	// sum2(20, 10)
	// ret, r := func3(20, 10)
	// fmt.Println(ret)
	// fmt.Println(r)
	// countStr()

	// arr := [...]int{1, 23, 5, 543}
	// arr1 := [...][3]int{
	// 	{1, 5, 8},
	// 	{6, 9, 21},
	// }
	// fmt.Println(arr)
	// fmt.Println(arr1)

	//切片
	// var s1 []int{1,23,5,73,432,3}
	//声明
	var s1 []int
	fmt.Println(s1)

	//make 声明并初始化
	var s2 = make([]int, 0, 10)

	fmt.Println(s2)
	fmt.Println(s2 == nil)
	// fmt.Printf("%T:"s1)
}
