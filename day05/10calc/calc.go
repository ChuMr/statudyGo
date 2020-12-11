package calc

import (
	"fmt"
)

func init() {
	fmt.Println("calc 初始化")
}

// Add 加法
func Add(x, y int) int {
	return x + y
}
