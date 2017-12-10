package main

import (
	"flag"
	"fmt"
)

func main() {
	act := flag.String("act", "producer", "Either: producer or consumer")
	//partition := flag.String("partition", "0",
	//	"Partition which the consumer program will be subscribing")

	flag.Parse()

	fmt.Printf("Welcome to go-microservice : %s\n\n", *act)

	switch *act {
	case "producer":
		mainProducer()
	case "consumer":
		//if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
		//	mainConsumer(int32(part32int))
		mainConsumer()
		//}
	}
}
