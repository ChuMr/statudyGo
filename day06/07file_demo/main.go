package main

import (
	"fmt"
	"os"
)

//获取文件对象的详细信息
func main() {
	fileObj, err := os.Open("./main.go")

	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("%T\n", fileObj)
	fmt.Printf("文件大小%dB \n", fileInfo.Size())
}
