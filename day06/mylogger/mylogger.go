package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// LogLevel 注释
type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAl
)

type logger struct {
	level LogLevel
}

func parseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "trace":
		return TRACE
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAl
	default:
		return DEBUG
	}
}

// NewLog 实例化日志类
func NewLog(level string) logger {
	LogLevel := parseLogLevel(level)
	return logger{
		level: LogLevel,
	}
}

func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime caller failed\n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

// NewLog 实例化日志类
func (l logger) myLog(DEBUG, msg string) {

	t := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", t.Format("2006-01-02 15:04:05"), DEBUG, fileName, funcName, lineNo, msg)
}

func (l logger) Debug(msg string) {
	if l.level <= DEBUG {
		l.myLog("DEBUG", msg)
	}
}

func (l logger) Info(msg string) {
	if l.level <= INFO {
		l.myLog("INFO", msg)
	}
}

func (l logger) Warning(msg string) {
	if l.level <= WARNING {
		l.myLog("WARNING", msg)
	}
}

func (l logger) Error(msg string) {
	if l.level <= ERROR {
		l.myLog("ERROR", msg)
	}
}

func (l logger) Fatal(msg string) {
	if l.level <= FATAl {
		l.myLog("FATAl", msg)
	}
}
