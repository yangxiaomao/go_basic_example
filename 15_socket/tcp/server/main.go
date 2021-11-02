package main

import (
	"bufio"
	"fmt"
	"gotest/01/15_socket/proto"
	"io"
	"net"
)

// TCP server端

// TCP 经过编码/解码解决黏包问题
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

// // TCP黏包测试
// func process(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	var buf [1024]byte
// 	for {
// 		n, err := reader.Read(buf[:])
// 		if err == io.EOF {
// 			fmt.Println("内容为空")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("read from client failed, err:", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到client发来的数据：", recvStr)
// 	}
// }

// func main() {
// 	listen, err := net.Listen("tcp", "127.0.0.1:30000")
// 	if err != nil {
// 		fmt.Println("listen failed, err:", err)
// 		return
// 	}
// 	defer listen.Close()
// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			fmt.Println("accept failed, err:", err)
// 			continue
// 		}
// 		go process(conn)
// 	}
// }

// // 处理函数
// func process(conn net.Conn) {
// 	defer conn.Close() // 关闭连接
// 	for {
// 		reader := bufio.NewReader(conn)
// 		var buf [128]byte
// 		n, err := reader.Read(buf[:]) // 读取数据
// 		if err != nil {
// 			fmt.Println("read form client failed, err:", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到client端发来的数据：", recvStr)
// 		conn.Write([]byte(recvStr))
// 	}
// }

// func main() {
// 	listen, err := net.Listen("tcp", "127.0.0.1:20000")
// 	if err != nil {
// 		fmt.Println("listen failed, err:", err)
// 		return
// 	}
// 	for {
// 		conn, err := listen.Accept() // 建立连接
// 		if err != nil {
// 			fmt.Println("accept failed, err:", err)
// 		}
// 		go process(conn) // 启动一个goroutine处理连接
// 	}
// }
