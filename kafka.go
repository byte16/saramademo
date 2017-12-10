package main

import (
	json "encoding/json"
	"fmt"
	"github.com/Shopify/sarama" //Kafka通讯库：Sarama
	"github.com/bsm/sarama-cluster"
	"os"
)

var (
	brokers       = []string{"192.168.1.108:9092"}
	topic         = "go-microservice-transactions"
	topics        = []string{topic}
	consumerGroup = "consumer-a"
)

func newKafkaProducerCfg() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaConsumerCfg() *cluster.Config {
	conf := cluster.NewConfig()
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Group.Mode = cluster.ConsumerModePartitions
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(brokers, newKafkaProducerCfg())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}
	return kafka
}

func newKafkaConsumer() *cluster.Consumer {
	consumer, err := cluster.NewConsumer(brokers, consumerGroup, topics, newKafkaConsumerCfg())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return consumer
}

func sendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n",
		partition, offset)

	return nil
}
