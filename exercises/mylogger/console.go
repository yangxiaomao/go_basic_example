package mylogger

import (
	"fmt"
	"time"
)

// 往终端中写日志相关内容

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.Level

}

func (l ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-01 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}

}

func (l ConsoleLogger) Debug(format string, arg ...interface{}) {
	l.log(DEBUG, format, arg...)
}

func (l ConsoleLogger) Info(format string, arg ...interface{}) {
	l.log(INFO, format, arg...)
}

func (l ConsoleLogger) Warning(format string, arg ...interface{}) {
	l.log(WARNING, format, arg...)
}

func (l ConsoleLogger) Error(format string, arg ...interface{}) {
	l.log(ERROR, format, arg...)
}

func (l ConsoleLogger) Fatal(format string, arg ...interface{}) {
	l.log(FATAL, format, arg...)
}
