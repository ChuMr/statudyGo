package main

import (
	"fmt"
	"sync"
	"time"
)

// 当读的操作远远大于写的操作的时候，读写锁的一定会比互斥锁高的

var x = 0
var lock sync.Mutex
var rwLock sync.RWMutex
var wg sync.WaitGroup

func lockRead() {
	defer wg.Done()
	rwLock.RLock()
	// lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	// lock.Unlock()
}

func lockWrite() {
	defer wg.Done()
	// lock.Lock()
	rwLock.Lock()
	x++
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()
	rwLock.Unlock()
}

func main() {
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		go lockWrite()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go lockRead()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(startTime))
}
