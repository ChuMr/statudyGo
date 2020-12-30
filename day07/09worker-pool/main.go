package main

import (
	"fmt"
	"time"
)

// worker-pool

// <-chan  只读的通道
// chan <- 只写的通道
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//开启3个goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	//5 个任务
	for j := 1; j <= 5; j++ {
		jobs <- j //jobs 写入
	}
	close(jobs)

	//输出结果 3个goroutine，去执行5个goroutine
	for a := 1; a <= 5; a++ {
		<-result
	}
}
