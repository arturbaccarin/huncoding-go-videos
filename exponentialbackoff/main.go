// https://youtu.be/6EVbp0Usqv0 7:30
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

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
