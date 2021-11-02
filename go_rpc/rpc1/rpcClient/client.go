package main

import (
	"fmt"
	"net/rpc"
)

type Params struct {
	Width,Height int
}

// 调用服务
func main(){
	// 1. 链接远程RPC服务
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
    if err != nil {
    	fmt.Println(err)
		return
	}
	// 2.调用远程方法
	// 2.1 求面积
	ret := 0
	err = rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.1 求周长
	ret2 := 0
	err = rp.Call("Rect.Perimeter", Params{50, 100}, &ret2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("矩形面积是：%v\n", ret)
	fmt.Printf("矩形周长是：%v\n", ret2)

}