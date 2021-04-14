package mylogger

import (
	"fmt"
	"time"
)

type Consolelogger struct {
	level LogLevel
}

// NewLog 实例化日志类
func NewLogConsoleLogger(level string) Consolelogger {
	LogLevel := parseLogLevel(level)
	return Consolelogger{
		level: LogLevel,
	}
}

// NewLog 实例化日志类

func (l Consolelogger) myLog(lv LogLevel, format string, a ...interface{}) {
	if lv >= l.level {
		msg := fmt.Sprintf(format, a...)
		t := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		logStr := unParseLogLevel(lv)
		fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", t.Format("2006-01-02 15:04:05"), logStr, fileName, funcName, lineNo, msg)
	}
}

func (l Consolelogger) Debug(format string, a ...interface{}) {
	l.myLog(DEBUG, format, a...)
}

func (l Consolelogger) Info(format string, a ...interface{}) {
	l.myLog(INFO, format, a...)
}

func (l Consolelogger) Warning(format string, a ...interface{}) {
	l.myLog(WARNING, format, a...)
}

func (l Consolelogger) Error(format string, a ...interface{}) {
	l.myLog(ERROR, format, a...)
}

func (l Consolelogger) Fatal(format string, a ...interface{}) {
	l.myLog(FATAl, format, a...)
}
