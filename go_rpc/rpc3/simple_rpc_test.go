package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

// 用于测试的结构体
// 字段首字母必须大写
type User struct {
	Name string
	Age int
}

// 用于测试的查询用户的方法
func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	user[0] = User{"zs", 20}
	user[1] = User{"ls", 21}
	user[2] = User{"ww", 22}
	// 模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("id %d not in user db", uid)
}

//  测试方法
func TestRPC(t *testing.T) {
	// 需要对interface{} 可能产生的类型进行注册
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务端
	srv := NewServer(addr)
	// 将方法注册到服务端
	srv.Register("queryUser", queryUser)
	// 服务端等待调用
	go srv.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	// 创建客户端
	cli := NewClient(conn)
	// 声明函数原型
	var query func(int)(User, error)
	cli.callRPC("queryUser", &query)
	// 得到查询结果
	u, err := query(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}