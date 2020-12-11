package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 1 往不同的位置记录日志
	// 2 日志分为5种级别
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("您输入的内容为:")
	s, _ := reader.ReadString('1')
	fmt.Printf("您输入的内容为:%s", s)
}
