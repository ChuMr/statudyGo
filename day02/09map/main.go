package main

import (
	"fmt"
)

//map
func main() {
	// var m map[string]int

	// // m = make(map[string]int, 10)
	// // fmt.Println(m)
	// // m["age"] = 20
	// // m["name"] = 5656
	// // val, ok := m["name"]
	// // fmt.Println(val)
	// // fmt.Println(ok)
	// // fmt.Println(m["name1"])

	// //按照指定的key

	// rand.Seed(time.Now().UnixNano())
	// var scoreMap = make(map[string]int, 200)
	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%02d", i)
	// 	value := rand.Intn(100)
	// 	scoreMap[key] = value
	// }

	// var keys = make([]string, 0, 200)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }

	// // fmt.Println(scoreMap)
	// fmt.Println(keys)

	var s1 = make([]map[int]string, 10, 10)
	fmt.Println(s1[0])
}
