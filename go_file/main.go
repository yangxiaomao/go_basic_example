package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 借助 io.Copy()实现一个拷贝的文件函数
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') // 注意是字符
		if err == io.EOF {
			// 退出之前将已读到的内容输出
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
	// // io.Copy()
	// _, err := CopyFile("./dst.txt", "./xx.txt")
	// if err != nil {
	// 	fmt.Println("copy file failed, err:", err)
	// 	return
	// }
	// fmt.Println("copy done!")
	// // ioutil.WriteFile
	// str := "hello 毛总"
	// err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	// if err != nil {
	// 	fmt.Println("write file failed, err:", err)
	// 	return
	// }
	// // bufio.NewWriter
	// file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println("open file failed, err:", err)
	// 	return
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 10; i++ {
	// 	writer.WriteString("hello毛总\n") // 将数据先写入缓存
	// }
	// writer.Flush() // 将缓存中的内容写入文件
	// // 写操作， Write和WriteString
	// file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println("open file failed, err:", err)
	// 	return
	// }
	// defer file.Close()
	// str := "hello 毛总"
	// file.Write([]byte(str))       // 写入字节切片数据
	// file.WriteString("hello 小王子") // 直接写入字符串数据
	// // ioutil读取整个文件
	// content, err := ioutil.ReadFile("./main.go")
	// if err != nil {
	// 	fmt.Println("read file failed, err:", err)
	// 	return
	// }
	// fmt.Println(string(content))

	// // file.Read()
	// // 只读方式打开当前目录下的main.go文件
	// file, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Println("open file failed!, err:", err)
	// 	return
	// }
	// defer file.Close()
	// // bufio读取文件
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n') // 注意是字符
	// 	if err == io.EOF {
	// 		if len(line) != 0 {
	// 			fmt.Println(line)
	// 		}
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed, err:", err)
	// 		return
	// 	}
	// 	fmt.Println(line)
	// }
	// // 使用Read方法读取数据
	// var content []byte
	// var tmp = make([]byte, 128)
	// for {
	// 	n, err := file.Read(tmp)
	// 	if err == io.EOF {
	// 		fmt.Println("文件读完了")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("read file failed, err:", err)
	// 		return
	// 	}
	// 	content = append(content, tmp[:n]...)
	// }

	// fmt.Println(string(content))
	// // 只读方式打开当前目录下的main.go文件
	// file, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Println("open file failed!, err:", err)
	// 	return
	// }
	// fmt.Println(file.Name())
	// // 关闭文件
	// file.Close()
}
