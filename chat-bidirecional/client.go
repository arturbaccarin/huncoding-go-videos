package main

import (
	"bdchat/chat"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer connection.Close()

	chatClient := chat.NewChatServiceClient(connection)

	stream, err := chatClient.Join(context.Background())
	if err != nil {
		log.Fatalf("failed to join: %v", err)
	}

	fmt.Println("Digite seue nome:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user := scanner.Text()

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("failed to recv: %v", err)
			}

			fmt.Printf("[%s] %s: %s", time.Unix(msg.Timestamp, 0).Format("15:04:05"), msg.User, msg.Text)
		}
	}()

	for scanner.Scan() {
		msg := &chat.Message{
			User:      user,
			Text:      scanner.Text(),
			Timestamp: time.Now().Unix(),
		}

		if err := stream.Send(msg); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
	}
}
