package main

import (
	"unaryrpc/client"
	"unaryrpc/server"
)

// protoc --go_out=. --go-grpc_out=. user.proto
func main() {
	go server.Run()
	client.Run()
}
