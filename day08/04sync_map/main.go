package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var wg sync.WaitGroup
var sm sync.Map

func get(key string) int {
	return m[key]
}

func set(key string, val int) {
	m[key] = val

}

// func main() {
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			set(key, n)
// 			fmt.Printf("k=:%v,v=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()
// }

func main() {
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// set(key, n)
			sm.Store(key, n)
			value, _ := sm.Load(key)
			fmt.Printf("k=:%v,v=%v\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
