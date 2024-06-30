package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go printUntilCancel(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func printUntilCancel(ctx context.Context) {
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel signed received")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(count)
			count++
		}
	}
}
