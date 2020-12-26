package mylogger

import (
	"fmt"
	"time"
)

type Conolelogger struct {
	level LogLevel
}

// NewLog 实例化日志类
func NewLog(level string) Conolelogger {
	LogLevel := parseLogLevel(level)
	return Conolelogger{
		level: LogLevel,
	}
}

// NewLog 实例化日志类
func (l Conolelogger) myLog(DEBUG, format string, a ...interface{}) {

	msg := fmt.Sprintf(format, a...)
	t := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", t.Format("2006-01-02 15:04:05"), DEBUG, fileName, funcName, lineNo, msg)
}

func (l Conolelogger) Debug(format string, a ...interface{}) {
	if l.level <= DEBUG {
		l.myLog("DEBUG", format, a...)
	}
}

func (l Conolelogger) Info(format string, a ...interface{}) {
	if l.level <= INFO {
		l.myLog("INFO", format, a...)
	}
}

func (l Conolelogger) Warning(format string, a ...interface{}) {
	if l.level <= WARNING {
		l.myLog("WARNING", format, a...)
	}
}

func (l Conolelogger) Error(format string, a ...interface{}) {
	if l.level <= ERROR {
		l.myLog("ERROR", format, a...)
	}
}

func (l Conolelogger) Fatal(format string, a ...interface{}) {
	if l.level <= FATAl {
		l.myLog("FATAl", format, a...)
	}
}
