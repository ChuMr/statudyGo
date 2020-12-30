package main

import "fmt"

// channel 操作
// 发送 ch1 <- 10
// 接收 x := <- ch1
// 关闭 close()

var c chan int

func main() {
	//slice map chan

	c = make(chan int, 10)
	fmt.Println(c)
	c <- 10
	fmt.Println(10)
	x := <-c
	fmt.Println(x)

}
