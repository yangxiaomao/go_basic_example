package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log_transfer/conf"
	"log_transfer/es"
	"log_transfer/kafka"
)

// log transfer
// 将日志数据从kafka取出来发往ES

func main(){
	// 0.加载配置文件
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config, err:%v\n", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)

	// 1.初始化ES
	// 1.2 初始化一个ES连接的client

	err = es.Init(cfg.EsCfg.Address)
	if err != nil {
		fmt.Printf("init ES client failed, err:%v\n", err)
		return
	}

	// 2.初始化kafka
	// 2.1连接kafka, 创建分区的消费者
	// 2.2每个分区的消费者分别取出数据 通过SendToEs（）将数据发送ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err:%v\n", err)
		return
	}
	select {

	}


	// 2. 从kafka中取日志

	// 3.发往ES
}
