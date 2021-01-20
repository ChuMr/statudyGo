package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1 与服务端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Printf("conn 127.0.0.1:10000 err%v", err)
	}

	//2 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {

		// var tmp [128]byte
		// n, err := conn.Read(tmp[:])
		// fmt.Println(n, "test")
		// if err != nil {
		// 	fmt.Printf("byte Read falied,err:%v", err)
		// }

		// fmt.Println(string(tmp[:n]))

		fmt.Print("请输入要发送的内容:")
		msg, _ := reader.ReadString('\n') //读取到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			conn.Close()
			break
		}
		conn.Write([]byte(msg))

	}
}
