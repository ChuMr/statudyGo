package main

import "fmt"

var x = 0

func f1() {
	for i := 0; i <= 50000; i++ {
		x++
	}
}
func main() {
	f1()
	f1()
	fmt.Println(x)
}
