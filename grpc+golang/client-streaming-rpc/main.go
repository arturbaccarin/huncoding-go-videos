package main

import (
	"clientstreaming/client"
	"clientstreaming/server"
)

func main() {
	go server.Run()
	client.Run()
}
