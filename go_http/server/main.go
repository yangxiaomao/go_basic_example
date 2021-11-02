package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 对应的Server端HeandlerFuncPost如下：
func getHandlerPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1.请求类型是application/x-www-form-urlencode时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

// 对应的Server端HandlerFuncGet如下：
func getHandlerGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}

// 测试服务端
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 毛总！")
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
	// http.HandleFunc("/post", getHandlerPost)
	// http.ListenAndServe("127.0.0.1:9000", nil)
}
