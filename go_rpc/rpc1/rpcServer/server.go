package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Rect struct {

}

// 声明参数结构体，字段首字母大写
type Params struct {
	// 长和宽
	Width, Height int
}

// 定义求矩形面积的方法
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// 定义求矩形周长的方法
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main(){
	// 1、注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 2.把服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务， 等待客户端调用求面积和周长的方法
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
