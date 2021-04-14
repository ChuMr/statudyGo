package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger ...
type FileLogger struct {
	level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *LogMsg
}

// LogMsg ...
type LogMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	lineNum   int
}

// NewFileLogger ...
func NewFileLogger(levelStr, fp, fn string, maxFile int64) *FileLogger {
	LogLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxFile,
		logChan:     make(chan *LogMsg, 50000),
	}
	err := fl.initFile() //打开文件
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed err:%v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+"err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed err:%v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	//开启后台goroutine 写日志

	// for i := 0; i <= 5; i++ {

	// }

	go f.writeLogBackground()

	return nil
}

func (f *FileLogger) fileClose() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat failed err:%v", err)
		return false
	}
	// fmt.Println(fileInfo.Size())
	// fmt.Println(f.maxFileSize)
	// fmt.Println("checkSize-------------------")
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割文件
	// 1 关闭当前文件
	// f.fileClose()

	// 2 rename 备份
	nowStr := time.Now().Format("20060102150405") // xx.log =>xx.log

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("err:%v", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name()) //拿到当前日志文件的完整路径
	newLogName := fmt.Sprintf("%s.back%s", logName, nowStr)

	file.Close()
	// fmt.Println(f.filePath)
	// fmt.Println(fileInfo.Name())
	// fmt.Println(logName)
	// fmt.Println(newLogName)
	// fmt.Println("test")
	err = os.Rename(logName, newLogName)
	// err = os.Rename(logName, )

	if err != nil {
		fmt.Printf("Rename file failed, err:%v\n", err)
	}
	// 3 打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
	}
	return fileObj, err
	// f.fileObj = fileObj
}

// NewLog 实例化日志类
func (f *FileLogger) myLog(lv LogLevel, format string, a ...interface{}) {
	if lv >= f.level {
		msg := fmt.Sprintf(format, a...)
		t := time.Now()
		funcName, fileName, lineNo := getInfo(3)

		logTmp := &LogMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: t.Format("2006-01-02 15:04:05"),
			lineNum:   lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			//把日志丢掉
		}

	}
}

func (f *FileLogger) writeLogBackground() {
	// 切割日志文件
	if f.checkSize(f.fileObj) {
		newFile, err := f.splitFile(f.fileObj)
		if err != nil {
			return
		}
		f.fileObj = newFile
	}

	for {

		select {
		case logTmp := <-f.logChan:
			logStr := unParseLogLevel(logTmp.level)
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d]%s\n", logTmp.timestamp, logStr, logTmp.fileName, logTmp.funcName, logTmp.lineNum, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				//如果日志级别大于 error级别，还需要在日志文件中记录
				if f.checkSize(f.errFileObj) {
					//切割日志文件
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						fmt.Println("return ?1?")
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}

	}
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.myLog(DEBUG, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.myLog(INFO, format, a...)
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.myLog(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.myLog(ERROR, format, a...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.myLog(FATAl, format, a...)
}
