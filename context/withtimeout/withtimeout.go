package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	count := 1
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
