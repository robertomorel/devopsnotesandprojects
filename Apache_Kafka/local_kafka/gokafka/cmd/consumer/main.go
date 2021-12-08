package main

import ( 
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap {
		"bootstrapserver": "gokafka_kafka_1:9092",
		"client_id": "goapp-consumer",
		"group.id": "goapp-group", // Podemos colocar vários consumers no mesmo grupo e cada um lê de uma partição diferente
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	// Fica lendo do pópico
	for {
		msg, err := c.ReadMessage(-1 /* Timeout */)
		if err == nil {
			fmt.Println(string(msg.value), msg.TopicPartition)
		}
	}
}