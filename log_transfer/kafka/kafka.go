package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log_transfer/es"
)


// 初始化kafka消费者，从kafka取数据发往ES
func Init(addrs []string, topic string) (error) {
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		fmt.Println("pc:", pc)
		fmt.Println(topic)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				// 直接发送ES
				ld := es.LogData{
					Topic: topic,
					Data:string(msg.Value),
				}
				fmt.Println(ld)
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
				es.SendToEsChan(&ld)
			}
		}(pc)
	}
	return err
}