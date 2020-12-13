package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["test"] = 1000
	fmt.Println(m1)
}
