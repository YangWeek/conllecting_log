package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

// kafka 消费者 如何从认定分区中取出数据
// 取出数据
func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("sarama consumer star failed, err is %v\n", err)
		return
	}
	// 拿到指topic下的全部的的分区
	partitions, err := consumer.Partitions("logagent")
	if err != nil {
		fmt.Printf("failed to get Partitions, err is %v\n", err)
		return
	}
	fmt.Println(partitions)

	for partition := range partitions {
		pc, err := consumer.ConsumePartition("logagent", int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("consumePartiton failed, error is %v\n", err)
			return
		}
		defer pc.AsyncClose() //
		// 异步消费
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, offest:%d, key:%v, val:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	select {}
}
