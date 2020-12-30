package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//goroutine 面试题
// goroutine 与线程的关系（os的线程 ）
//
//随机数
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r := rand.Int()
		r1 := rand.Intn(10)
		fmt.Println(r, r1)
	}
}

func f1(i int) {

	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()

}
