package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func fileWrite1() {
	fileObj, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("file open failed %v", err)
	}

	fileObj.Write([]byte("say hello word\n")) //字符串转换为切片
	fileObj.WriteString("上海\n")
}

//bufio
func fileWrite2() {
	fileObj, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("file open failed %v", err)
	}
	w := bufio.NewWriter(fileObj)
	w.Write([]byte("hello word\n"))
	w.WriteString("中国\n")
	w.Flush()
}

func fileWrite3() {
	err := ioutil.WriteFile("test.txt", []byte("fileWrite3\n"), 0666)
	if err != nil {
		fmt.Printf("file open failed %v", err)
		return
	}
}

func readFromFile(filePath string) (string, error) {
	fileObj, err := os.Open(filePath)
	if err != nil {
		// fmt.Printf("open file failed ,err:%v", err)

		return "", err
	}
	defer fileObj.Close()

	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if n <= 0 {
			return "", nil
		}
		if err != nil {
			// fmt.Printf("read file failed ,err:%v", err)
			return "", err
		}
		return string(tmp[:n]), nil
		// fmt.Printf("读了%d 个字节", n)
		// fmt.Println(string(tmp[:n]))
	}
}

func fileCopy(file string, copyFile string) {

	s, err := readFromFile(file)
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return
	}

	fileObj, err := os.OpenFile(copyFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open copyfile failed err:%v", err)
		return
	}
	fileObj.WriteString(s)
	// fileObj.Write([]byte("say hello word\n")) //字符串转换为切片
	// fileObj.WriteString("上海\n")
}

//作业copy file
//文件写入
func main() {
	// fileWrite1()
	// fileWrite2()
	// fileWrite3()
	fileCopy("test.txt", "copy.txt")
}
