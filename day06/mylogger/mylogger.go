package mylogger

import (
	"fmt"
	"time"
)

const (
	DEBUG Loglevel = iota
	INFO
	Warning
	WARNING
	Error
	FATAl
)

type logger struct {
	level
}

// NewLog 实例化日志类
func NewLog(l logger.Loglevel) logger {
	return logger{
		level: level,
	}
}

func (l logger) Debug(msg string) {
	t := time.Now()
	fmt.Printf("[%v] %s\n", t.Format("2006-01-02 15:04:05"), msg)
}

func (l logger) Info(msg string) {
	t := time.Now()
	fmt.Printf("[%v] %s\n", t.Format("2006-01-02 15:04:05"), msg)
}

func (l logger) Warning(msg string) {
	t := time.Now()
	fmt.Printf("[%v] %s\n", t.Format("2006-01-02 15:04:05"), msg)
}

func (l logger) Error(msg string) {
	t := time.Now()
	fmt.Printf("[%v] %s\n", t.Format("2006-01-02 15:04:05"), msg)
}

func (l logger) Fatal(msg string) {
	t := time.Now()
	fmt.Printf("[%v] %s\n", t.Format("2006-01-02 15:04:05"), msg)
}
