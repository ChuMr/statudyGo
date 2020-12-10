package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	dis = make(map[string]int, len(users))
)

func main() {
	//50 枚金币
	//字符串里的每个字母是字符 字符
	for _, user := range users {
		// r := []rune(user)
		// fmt.Println(user)
		// fmt.Println(r)
		for _, val := range user {
			fmt.Printf("%q", val)
			fmt.Println(string(val))
			// fmt.Println(val)
			fmt.Printf("%T\n", val)
			// fmt.Printf("%v\n", val)
			// fmt.Printf("%s\n", val)

			// if val == 'E' || val == 'e' {
			// 	// fmt.Printf("%v\n", val)
			// 	// fmt.Println(val)
			// 	// fmt.Printf("%v\n", val)
			// 	// fmt.Println(user)
			// }
		}
	}
}
