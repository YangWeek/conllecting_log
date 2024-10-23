package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// 连接 kafka的 客户端
var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage // 生产者消息
)

// Message 发送到kafka的message
type Message struct {
	Data  string
	Topic string
}

func InitKafka(addrs []string, chansize int64) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		logrus.Error("producer closed, err:", err)
		return
	}
	msgChan = make(chan *sarama.ProducerMessage, chansize)
	go SendMsg()
	return nil
}

func SendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warnf("send message failed, err is %v\n", err)
				return
			}
			logrus.Infof("send msg to kafka success, pid is %v,offset is %v\n", pid, offset)
		}
	}
}

func SetKafkaChanMes(msg *sarama.ProducerMessage) {
	msgChan <- msg
	return
}

// SendLog 往msgChan发送消息的函数
func SendLog(msg1 *Message) (err error) {
	//// 包装一下这个Message
	//var data []byte
	//err = json.Unmarshal(msg1.Data, data)
	//if err != nil {
	//	logrus.Errorf("sys data unmarshal failed, err is %v\n", err)
	//}

	msg := &sarama.ProducerMessage{
		Topic: msg1.Topic,
		Value: sarama.StringEncoder(msg1.Data), // 这个就行
	}
	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("msgChan is full")
	}
	return
}

//func sendKafka() {
//	for msg := range msgChan {
//		kafkaMsg := &sarama.ProducerMessage{}
//		kafkaMsg.Topic = msg.Topic
//		kafkaMsg.Value = sarama.StringEncoder(msg.Data)
//		pid, offset, err := client.SendMessage(kafkaMsg)
//		if err != nil {
//			logrus.Warnf("send msg failed, err:%v\n", err)
//			continue
//		}
//		logrus.Infof("send msg success, pid:%v offset:%v\n", pid, offset)
//	}
//}
