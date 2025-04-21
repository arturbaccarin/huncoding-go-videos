package main

import (
	"bdchat/chat"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type chatServer struct {
	chat.UnimplementedChatServiceServer
	mu       sync.Mutex
	clients  map[chat.ChatService_JoinServer]bool
	messages chan *chat.Message
}

func newServer() *chatServer {
	return &chatServer{
		clients:  make(map[chat.ChatService_JoinServer]bool),
		messages: make(chan *chat.Message),
	}
}

func (s *chatServer) Join(stream chat.ChatService_JoinServer) error {

	s.mu.Lock()
	s.clients[stream] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				return
			}

			s.messages <- msg
		}
	}()

	for message := range s.messages {
		s.mu.Lock()
		for client := range s.clients {
			if err := client.Send(message); err != nil {
				fmt.Printf("failed to send message: %v", err)
			}
		}
		s.mu.Unlock()
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, newServer())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
