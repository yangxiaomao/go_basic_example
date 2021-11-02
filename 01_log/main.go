package main

import (
	"gotest/01/01_log/logs"
)

func main() {
	logs.NewLog("success", "log_time_", "测试日志内容")
}
