package main

import "fmt"

func adder() func(int) int {
	var x = 100
	return func(y int) int {
		fmt.Println(y, "test")
		x += y
		return x
	}
}
func main() {
	ret := adder()
	fmt.Printf("%T \n", ret)
	// fmt.Println("ret")
	// ret1 := ret(200)
	// fmt.Println(ret1)
}
