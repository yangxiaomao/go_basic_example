package main

import (
	"fmt"
	"gotest/01/15_socket/proto"
	"net"
)

// TCP client客户端

// TCP 经过编码解决黏包问题
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}

// // TCP 黏包测试
// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:30000")
// 	if err != nil {
// 		fmt.Println("dial failed, err:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for i := 0; i < 20; i++ {
// 		msg := `Hello, Hello, How are you?`
// 		fmt.Println(msg)
// 		conn.Write([]byte(msg))
// 	}
// }

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:20000")
// 	if err != nil {
// 		fmt.Println("err :", err)
// 		return
// 	}
// 	defer conn.Close() // 关闭连接
// 	inputReader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := inputReader.ReadString('\n') // 读取用户输入
// 		inputInfo := strings.Trim(input, "\r\n")
// 		if strings.ToUpper(inputInfo) == "Q" { // 如果输入Q就退出
// 			return
// 		}

// 		_, err = conn.Write([]byte(inputInfo)) // 发送数据
// 		if err != nil {
// 			return
// 		}
// 		buf := [512]byte{}
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Println("recv failed, err:", err)
// 			return
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }
