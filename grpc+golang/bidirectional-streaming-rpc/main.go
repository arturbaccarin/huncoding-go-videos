// https://youtu.be/gghOIrJrRTY

package main

import (
	"birpc/client"
	"birpc/server"
)

func main() {
	go server.RunServer()
	go client.RunClient()

	ch := make(chan int)
	<-ch
}
