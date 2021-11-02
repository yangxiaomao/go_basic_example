package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 声明接收的参数结构体
type ArithRequest struct {
	A,B int
}

// 声明返回客户端参数的结构体
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}


// 调用服务
func main(){
	// 连接远程的RPC
	//conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	req := ArithRequest{9, 2}
	var res ArithResponse
	// 调用乘法
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
	// 调用余数和商
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d / %d, 商 = %d, 余数 = %d\n", req.A, req.B, res.Quo, res.Rem)
}
