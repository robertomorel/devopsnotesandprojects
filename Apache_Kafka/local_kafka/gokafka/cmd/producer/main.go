package main

import ( 
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main()  {
	fmt.Println("main")

	/*
		Channel são eventos do Go que fazem comunicações de funções em contextos diferentes 
	*/
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()

	// Se a mensagem tiver sem key, cada hora ela cai em uma partição diferente
	// Publish(msg: "mensagem de teste", topic: "teste", producer, key: nil, deliveryChan)

	// Utilizando a key, após o primeiro envio, que o Kafka seleciona a partição, todos os outros envios que usam a mesma key irá para a mesma partição 
	Publish(msg: "mensagem de teste", topic: "teste", producer, []byte("transferencia"), deliveryChan)

	go DeliveryReport(deliveryChan) // "go" command vai jogar para outra thread de forma assíncrona
	producer.Flush(timeoutMs: 1000)

	// // Criando um evento para aguardar o deliveryChan ser preenchido
	// // Trava o programa enquanto não receber a mensagem
	// e := <- deliveryChan
	// // Pega a mensagem do evento
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("Erro ao enviar")
	// } else {
	// 	fmt.Println("Mensagem enviada: ", msg.TopicPartition)
	// }

	// // Esperar a mensagem ser entregue, ou determinado tempo
	// producer.Flush(timeoutMs: 1000)
}

// Criando novo Producer
func NewKafkaProducer() *kafka.Producer {
	// Criando mapa de configuração
	configMap := &kafka.ConfigMap {
		"bootstrapserver": "gokafka_kafka_1:9092"
		"delivery.timeout.ms": "0", // Tempo máximo de entrega de uma mensagem (recebimento do retorno da entrega) 
		"acks": "all", // Quem vai sincronizar a mensagem
		"enable.idempotence": "false", //
	}

	// Criando novo Producer
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p;
}

// Publicando uma mensagem específica
func Publish(msg string, topic string, producer *kafka.Producer, key []byte deliveryChan chan kafka.Event) error  {
	msg := &kafka.Message {
		value: []byte(msg), // Transforma mensagem em arr de bytes
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: key,
	}

	// Depois da publicação da mensagem, o resultado dela, irá para o canal deliveryChan
	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return err
}

// 
func DeliveryReport(deliveryChan chan, kafka.Event) {
	for e := deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada", ev.TopicPartition)
			}
		}
	}
}