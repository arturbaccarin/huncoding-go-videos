package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func processTask(id int) error {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	if rand.Float32() < 0.2 {
		return fmt.Errorf("task %d failed after %v", id, delay)
	}

	fmt.Printf("task %d completed after %v\n", id, delay)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxConcurrency := 5
	semaphore := make(chan struct{}, maxConcurrency)

	var g errgroup.Group
	var mu sync.Mutex
	totalTasks := 20
	completedTasks := 0

	for i := 1; i <= totalTasks; i++ {
		taskID := i
		g.Go(func() error {
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				err := processTask(taskID)
				if err != nil {
					cancel()
					return err
				}

				mu.Lock()
				completedTasks++
				mu.Unlock()

				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("All tasks completed. Completed tasks: %d\n", completedTasks)
	}

	fmt.Printf("Total tasks: %d\n", totalTasks)
}
