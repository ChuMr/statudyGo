package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime caller
func main() {

	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime caller failed\n")
	}
	funcname := runtime.FuncForPC(pc).Name()
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(path.Base(file))
	fmt.Println(funcname)
}
