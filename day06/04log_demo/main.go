package main

// 日志库需求
// 1 生产写文件，测试打印终端
// 2 log的5种级别
/*
   1 Debug
   2 Info
   3 Warning
   4 Error
   5 fatal
*/
// 3 日志要有个开关
// 4 完整的日志记录要包含有时间、行号、文件名、日志级别、日志信息
// 5 日志文件要切割

func main() {
	log := NewLog()
	//go env -w GO111MODULE auto
	// fmt.Println("test")
	log.Output("这是一条测试日志")
}
