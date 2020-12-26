package main

import (
	"fmt"
	"time"

	"qian.zhang.com/studygo/day06/mylogger"
)

//	"zhangqian.com/day06/mylogger"

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

	fmt.Println("test")
	log := mylogger.NewLog("error")

	// log.myLog("这是一条DEBUG日志")
	for {
		log.Debug("这是一条DEBUG日志")

		log.Info("这是一条Info日志")

		log.Warning("这是一条Warning日志")

		log.Error("这是一条Error日志,id:%d,name:%s", 10000, "张三")

		log.Fatal("这是一条Fatal日志")

		time.Sleep(5 * time.Second)
	}
}
