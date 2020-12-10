package main

import "fmt"

//切片
func main() {

	// s1 := []int{1, 2, 3}
	// s2 := []string{"平山村"}

	// // fmt.Println(s1)
	// // fmt.Println(s2)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

	// fmt.Printf("len(s1):%d cap(s2):%d\n", len(s1), cap(s1))
	// fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//由数组切割成数组
	//切片的容量就是地城数组的容量
	// s3 := []int{1, 2, 3, 24, 23, 656, 76, 8}

	// s10 := s3[0:5] //len 5 cap 8
	// s11 := s3[2:]  //len 3 cap 3
	// // s12 := s3[:]   //len 3 cap 3
	// s12 := s11[2:3]
	// fmt.Printf("len(s10):%d cap(s10):%d\n", len(s10), cap(s10))
	// fmt.Printf("len(s11):%d cap(s11):%d\n", len(s11), cap(s11))
	// fmt.Printf("len(s12):%d cap(s12):%d\n", len(s12), cap(s12))

	// s3[4] = 230
	// fmt.Println(s10)
	// fmt.Println(s11)
	// fmt.Println(s12)
	//切片是传引用类型

	//make 函数构造切片
	// var slic = []int{}
	// var arr = [...]int{}
	// fmt.Println(slic)
	// fmt.Println(arr)

	// s1 := []string{"上海", "北京", "深圳"}
	// s2 := s1
	// var s3 = make([]string, 3)
	// copy(s3, s1)
	// fmt.Println(s1, s2, s3)
	// s1[2] = "广州"
	// fmt.Println(s1, s2, s3)

	// s1 = append(s1, s2...)
	// fmt.Println(s1)

	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	s1 = append(s1[:1], s1[2:]...)
	// fmt.Println(s1)
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(x1)

}
