package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello word", i)
}

func main() {
	// for i := 0; i < 100; i++ {
	// 	go hello(i)
	// }

	for i := 0; i < 100; i++ {
		func() {
			fmt.Println("hello word", i)
		}()
	}

	fmt.Println("main")
	time.Sleep(time.Second)
}
