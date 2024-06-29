package main

import (
	"fmt"
	"time"
)

func f1(c <-chan int) {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("Waiting...")
		fmt.Println(<-c)
		fmt.Println("Finished!")
	}
}

func f2(c chan<- int) {
	c <- 42
}

func main() {
	c := make(chan int)

	go f1(c)

	time.Sleep(3 * time.Second)

	println("sending 1")
	f2(c)
	println("sended 1")

	println("sending 2")
	f2(c)
	println("sended 2")

	time.Sleep(20 * time.Second)
}
