package main

import (
	"context"
	"fmt"
	"time"
)

func getValueFromChannel(ch chan<- string, duration time.Duration) {
	time.Sleep(duration)

	ch <- "Terminamos por aqui"
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	ch := make(chan string)
	go getValueFromChannel(ch, 2*time.Second)

	ch2 := make(chan string)
	go getValueFromChannel(ch2, 2*time.Second)

	ch3 := make(chan string)
	go getValueFromChannel(ch3, 2*time.Second)

	select {
	case returnValue := <-ch:
		fmt.Println(returnValue)
	case returnValue := <-ch2:
		fmt.Println(returnValue)
	case returnValue := <-ch3:
		fmt.Println(returnValue)
	case <-ctx.Done():
		fmt.Println("TEMPO MÀXIMO ENCERRADO")
	}
}
