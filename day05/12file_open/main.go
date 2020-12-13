package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

// ReadFromFile 打开文件，读文件的三种方法
func ReadFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed ,err:%v", err)
	}
	defer fileObj.Close()

	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if n <= 0 {
			return
		}
		if err != nil {
			fmt.Printf("read file failed ,err:%v", err)
		}
		fmt.Printf("读了%d 个字节", n)
		fmt.Println(string(tmp[:n]))
	}
}

//bufio
func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed ,err:%v", err)
		return
	}
	defer fileObj.Close()
	read := bufio.NewReader(fileObj)
	for {
		line, err := read.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed ,err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {

	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("ioutil readFile failed %v", err)
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFileByBufio()
	readFromFileByIoutil()
}
