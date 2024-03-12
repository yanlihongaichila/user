package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func SendMessage(message string, topic string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic:     topic,                       //包含了消息的主题
		Partition: int32(0),                    //
		Key:       sarama.StringEncoder("key"), //
		Offset:    1,
	}
	//fmt.Println(topic)

	//valueByte, err := json.Marshal(message)
	//value := string(valueByte)
	//fmt.Println(string(valueByte))
	//if err != nil {
	//	return err
	//}
	msgType := topic
	msg.Topic = msgType
	//将字符串转换为字节数组
	msg.Value = sarama.ByteEncoder([]byte(message))
	//msgStr, err := json.Marshal(msg)
	//if err != nil {
	//	return
	//}
	//fmt.Println(value)
	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	_, _, err = producer.SendMessage(msg)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1232312312")
	//fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	return
}

func main() {
	SendMessage("test", "test")
}
