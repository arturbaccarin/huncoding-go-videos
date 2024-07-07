package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: "meu-topico",
		Value: sarama.StringEncoder("enviando via go"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada: %d - %d \n", partition, offset)

}
