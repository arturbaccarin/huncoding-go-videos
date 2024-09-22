package main

import (
	"serverstreaming/client"
	"serverstreaming/server"
)

func main() {
	go server.Run()
	client.Run()
}
