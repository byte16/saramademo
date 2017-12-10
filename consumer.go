package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bsm/sarama-cluster"
	"os"
)

func mainConsumer() {
	kafka := newKafkaConsumer()
	defer kafka.Close()

	quit := make(chan bool, 1)

	go func() {
		for {
			select {
			// Partitions returns the read channels for individual partitions of this broker.
			//
			// This will channel will only return if Config.Group.Mode option is set to
			// ConsumerModePartitions.
			//
			// The Partitions() channel must be listened to for the life of this consumer;
			// when a rebalance happens old partitions will be closed (naturally come to
			// completion) and new ones will be emitted. The returned channel will only close
			// when the consumer is completely shut down.
			case part, ok := <-kafka.Partitions():
				if !ok {
					return
				}

				// start a separate goroutine to consume messages
				go consumeEvents(kafka, part)
			case <-quit:
				return
			}
		}
	}()

	fmt.Println("Press [enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	quit <- true
	fmt.Println("Terminating...")
}

func consumeEvents(kafka *cluster.Consumer, consumer cluster.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	var bankAccount *BankAccount
	var err error

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if !ok {
				fmt.Println("msg channel closed")
				return
			}
			kafka.MarkOffset(msg, "")
			msgVal = msg.Value
			//
			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))

				switch logType {
				case "CreateEvent":
					event := new(CreateEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						bankAccount, err = event.Process()
					}
				default:
					fmt.Println("Unknown command: ", logType)
				}

				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Printf("%+v\n\n", *bankAccount)
				}
			}

		}
	}
}
