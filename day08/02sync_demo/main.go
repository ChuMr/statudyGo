package main

import (
	"fmt"
	"sync"
)

// 互斥锁
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func f1() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go f1()
	go f1()
	wg.Wait()
	fmt.Println(x)
}
