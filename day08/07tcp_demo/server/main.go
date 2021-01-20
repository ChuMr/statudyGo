package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1 本地端口启动服务

	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Printf("listen tcp 127.0.0.1 falied,err:%v", err)
	}

	for {
		//2 等待链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept conn falied,err:%v", err)
		}

		//3 与客户端通信
		go processConn(conn)
	}

}

func processConn(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("byte Read falied,err:%v", err)
		}

		fmt.Println(string(tmp[:n]))

		fmt.Print("请回复:")
		msg, _ := reader.ReadString('\n') //读取到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}
