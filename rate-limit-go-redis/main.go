//https://youtu.be/P1ZK1F8Xs9M

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type RateLimiter struct {
	client  *redis.Client
	limit   int
	window  time.Duration
	context context.Context
}

func NewRateLimiter(client *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		client:  client,
		limit:   limit,
		window:  window,
		context: context.Background(),
	}
}

func (rl *RateLimiter) Allow(key string) bool {
	return false
}

func rateLimiter(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
	})
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	rateLimiter := NewRateLimiter(client, 10, 1*time.Minute)

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n")
	})

	http.ListenAndServe(":8080", nil)
}
