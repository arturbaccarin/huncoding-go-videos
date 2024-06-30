package withdeadline

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	printUntilCancel(ctx)
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
