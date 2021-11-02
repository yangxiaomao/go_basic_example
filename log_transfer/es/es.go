package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"context"
	"time"
)

type LogData struct {
	Topic string `json:"topic"`
	Data string `json:"data"`
}

// 初始化全局clinet
var (
	client *elastic.Client
	ch = make(chan *LogData, 100000)
)

// 初始化ES,准备接收kafka那边发来的数据

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://"+address
	}
	//fmt.Println("address:",address)
	// 1.初始化连接，得到一个client
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return
	}

	fmt.Println("connect to es success")
	go sendToEs()
	return
}

// SendToEs 发送数据到ES
func SendToEsChan(msg *LogData){
	ch <- msg
}

func sendToEs(){
	for {
		select {
		case msg := <-ch:
			// 链式操作
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
			default:
				time.Sleep(time.Second)
		}
	}

}
