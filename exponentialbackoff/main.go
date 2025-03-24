// https://youtu.be/6EVbp0Usqv0 7:30
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/*
cenário 2 - jitter
*/

func performRequest(ctx context.Context, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if rand.Intn(3) != 0 {
			return errors.New("some error")
		}

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		fmt.Println("Request performed successfully")

		return nil
	}
}

func retryWithBackoff(ctx context.Context, url string, maxRetries int, baseDelay time.Duration, maxTimeout time.Duration) error {
	retryDelay := time.Millisecond * 100

	for attempt := 1; attempt < maxRetries; attempt++ {
		err := performRequest(ctx, url)
		if err == nil {
			fmt.Printf("Request performed successfully after %d attempts\n", attempt)
			return nil
		}

		jitter := time.Duration(rand.Int63n(int64(retryDelay)))
		sleepTime := retryDelay + jitter/2
		fmt.Printf("Retrying after %s (attempt %d)\n", sleepTime, attempt)

		select {
		case <-time.After(sleepTime):
			retryDelay *= 2
			if retryDelay > maxTimeout {
				retryDelay = maxTimeout
			}
		case <-ctx.Done():
			return fmt.Errorf("request failed after %d attempts: %w", attempt, ctx.Err())
		}
	}

	fmt.Println("Request failed after", maxRetries, "attempts")
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	url := "https://jsonplaceholder.typicode.com/posts/1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := retryWithBackoff(ctx, url, 7, time.Millisecond*100, time.Second*2)
	if err != nil {
		fmt.Println("Request failed:", err)
	}
}

/*
cenário 1

func performRequest(url string) error {
	if rand.Intn(2) == 0 {
		return errors.New("some error")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	fmt.Println("Request performed successfully")

	return nil
}

func retryWithBackoff(url string, maxRetries int) {
	retryDelay := time.Millisecond * 100

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := performRequest(url)
		if err == nil {
			return
		}

		fmt.Println("Retrying after", retryDelay)
		time.Sleep(retryDelay)

		retryDelay *= 2
	}

	fmt.Println("Request failed after", maxRetries, "attempts")
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// retentativas de forma exporadicas com backoff
	retryWithBackoff(url, 5)
}
*/
