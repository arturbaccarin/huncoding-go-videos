package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(
		"meu-topico",
		0,
		sarama.OffsetOldest,
	)
	if err != nil {
		panic(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	messages := partitionConsumer.Messages()
	for {
		select {
		case <-sigCh:
			fmt.Println("Saindo da aplicação")
			return

		case msg := <-messages:
			fmt.Printf("Mensagem recebida: %s \n", string(msg.Value))
		}
	}
}
