package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	brokers := "localhost:9092"
	groupID := "your_consumer_group"
	topic := "mytopic"

	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		event := consumer.Poll(100)
		switch ev := event.(type) {
		case *kafka.Message:
			fmt.Printf("Received message on topic %s: %s\n", ev.TopicPartition, string(ev.Value))
		case kafka.Error:
			fmt.Printf("Error: %v\n", ev)
		}
	}
}
