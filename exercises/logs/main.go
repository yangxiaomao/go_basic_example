package main

import (
	mylogger "gotest/01/exercises/mylogger"
)

var myLog mylogger.Logger

// 测试我们自己写的日志库
func main() {
	// myLog := mylogger.NewLog("debug")
	myLog := mylogger.NewFileLogger("Info", "./", "maozong.log", 10*1024*1024)
	// myLog = mylogger.NewConsoleLogger("Info")
	for {
		id := 10001
		name := "毛总"
		myLog.Debug("这是一条Debug日志,id：%d, name:%s", id, name)
		myLog.Info("这是一条Info日志")
		myLog.Warning("这是一条Warning日志")
		myLog.Error("这是一条Error日志")
		myLog.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
