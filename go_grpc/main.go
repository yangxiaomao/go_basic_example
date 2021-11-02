package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	firstpb "go_grpc/src/first"
	"io/ioutil"
	"log"
)

func main() {
	pm := NewPersonMessage()
	//
	//_ = writeToFile("person.bin", pm)
	//pm2 := &firstpb.PersonMessage{}
	//_ = readFromFile("person.bin", pm2)
	//fmt.Println(pm2)

	pmStr := toJson(pm)
	fmt.Println(pmStr)

	pm3 := &firstpb.PersonMessage{}
	_ = fromJson(pmStr, pm3)
	fmt.Println("pb struct:", pm3)
}

func fromJson(in string, pb proto.Message) error {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("读取JSON时发生错误", err.Error())
	}
	return nil
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{Indent: "    "}

	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("转化为JSON时发生错误", err.Error())
	}
	return str
}

func readFromFile(fileName string, pb proto.Message) error {
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("读取文件时发生错误", err.Error())
	}

	err = proto.Unmarshal(dataBytes, pb)
	if err != nil {
		log.Fatalln("转化为struct的时候发生错误", err.Error())

	}
	return nil
}


func writeToFile(fileName string, pb proto.Message) error {
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("无法序列化到bytes", err.Error())
	}

	if err := ioutil.WriteFile(fileName, dataBytes, 0644); err != nil {
		log.Fatalln("无法写入到文件", err.Error())
	}
	log.Println("成功写入到文件")
	return nil
}

func NewPersonMessage() *firstpb.PersonMessage {
	pm := firstpb.PersonMessage{
		Id:           123,
		IsAdult:      true,
		Name:         "Dave",
		LuckyNumbers: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(pm)

	pm.Name = "Nick"

	fmt.Println(pm)

	fmt.Printf("The Id is: %d\n", pm.GetId())
	return &pm
}
