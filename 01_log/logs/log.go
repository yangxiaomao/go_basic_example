package logs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//PathExists 判断文件夹是否存在
func isPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}

func NewLog(clogFileStringName string, clogPrefix string, clogContent string) {
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	day := strconv.Itoa(time.Now().Day())
	var cLogFile string
	if clogFileStringName == "" {
		cLogFile = "./runtime/log/" + year + "/" + month + "/" + day + "/"
	} else {
		cLogFile = "./runtime/" + clogFileStringName + "/log/" + year + "/" + month + "/" + day + "/"
	}

	if ok, _ := isPathExist(cLogFile); ok {
		logFile, err := os.OpenFile(cLogFile+day+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if nil != err {
			fmt.Println(err.Error())

			// panic(err)
		} else {
			//创建一个Logger
			//参数1：日志写入目的地
			//参数2：每条日志的前缀
			//参数3：日志属性
			loger := log.New(logFile, clogPrefix, log.Ldate|log.Ltime|log.Lshortfile)

			//SetFlags设置输出选项
			loger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

			loger.Println(clogContent)

		}
		logFile.Close()
	} else {
		fmt.Println("文件不存在")
	}

}
