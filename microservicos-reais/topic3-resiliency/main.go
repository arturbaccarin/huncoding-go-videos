package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

const (
	rateLimitRequestsPerSecond = 5
	bulkheadMaxConcurrent      = 3
)

var limiter = rate.NewLimiter(rateLimitRequestsPerSecond, 1)

var bulkhead = semaphore.NewWeighted(bulkheadMaxConcurrent)

func unstableService(ctx context.Context) error {

	if rand.Float32() < 0.5 {
		return errors.New("unstable service")
	}

	time.Sleep(200 * time.Millisecond)
	return nil
}

func executeWithRetry(ctx context.Context) error {
	retryPolicy := backoff.NewExponentialBackOff()
	retryPolicy.MaxElapsedTime = 5 * time.Second

	operation := func() error {
		if err := limiter.Wait(ctx); err != nil {
			return fmt.Errorf("rate limit exceeded: %w", err)
		}

		if err := bulkhead.Acquire(ctx, 1); err != nil {
			return fmt.Errorf("bulkhead exceeded: %w", err)
		}
		defer bulkhead.Release(1)

		if err := unstableService(ctx); err != nil {
			return fmt.Errorf("unstable service: %w", err)
		}

		return nil
	}

	return backoff.Retry(operation, retryPolicy)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := executeWithRetry(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("falha ao processar: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/process", handler)
	port := 8080
	log.Printf("Listening on http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
