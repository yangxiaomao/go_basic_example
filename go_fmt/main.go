package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	bufioDemo()
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// // fmt.Scanln
	// fmt.Scanln(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// Scanf
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	// // 获取输入
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	// // 整型
	// n := 65
	// fmt.Printf("%b\n", n) // 二进制
	// fmt.Printf("%c\n", n) // 该值对应的unicode码值
	// fmt.Printf("%d\n", n) // 十进制
	// fmt.Printf("%o\n", n) // 八进制
	// fmt.Printf("%x\n", n) // 十六进制，使用a-f
	// fmt.Printf("%X\n", n) // 十六进制，使用A-F
	// // 格式化占位符
	// fmt.Printf("%v\n", 100)
	// fmt.Printf("%v\n", false)
	// o := struct{ name string }{"小王子"}
	// fmt.Printf("%v\n", o)
	// fmt.Printf("%#v\n", o)
	// fmt.Printf("%T\n", o)
	// fmt.Printf("100%%\n")

	// // Errorf
	// // err := fmt.Errorf("这是一个错误")
	// e := errors.New("原始错误e")
	// w := fmt.Errorf("Wrap了一个错误%w", e)
	// fmt.Println(w)

	// // Sprint
	// s1 := fmt.Sprint("通州毛总")
	// name := "通州毛总"
	// age := 18
	// s2 := fmt.Sprintf("name:%s, age:%d", name, age)
	// s3 := fmt.Sprintln("通州毛总")
	// fmt.Println(s1, s2, s3)

	// // 向标准输出写入内容
	// fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	// fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("打开文件出错.err:", err)
	// 	return
	// }
	// name := "通州毛总"
	// // 向打开的文件句柄中写入内容
	// fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
}

// func main() {
// 	fmt.Print("在终端打印该信息。")
// 	name := "通州毛总"
// 	fmt.Printf("我是：%s\n", name)
// 	fmt.Println("在终端打印单独一行显示")
// }
