package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct {

}

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

// 乘法运算
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) error{
	res.Pro = req.A * req.B
	return nil
}

// 商和余数
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) error{
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 商
	res.Quo = req.A / req.B
	// 余数
	res.Rem = req.A % req.B
	return nil
}

//func main(){
//	// 注册服务
//	rpc.Register(new(Arith))
//	// 采用http作为rpc载体
//	rpc.HandleHTTP()
//	err := http.ListenAndServe(":8081", nil)
//	if err != nil {
//		fmt.Println(err)
//	}
//}


// jsonRPC 编码方式
func main(){
	// 注册服务
	rpc.Register(new(Arith))
	// 监听服务
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 循环监听服务
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		// 协程
		go func (conn net.Conn){
			fmt.Println("new Client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}