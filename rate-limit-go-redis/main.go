//https://youtu.be/P1ZK1F8Xs9M

package main

import (
	"fmt"
	"net"
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
	pipe := rl.client.TxPipeline()

	incr := pipe.Incr(rl.context, key)
	pipe.Expire(rl.context, key, rl.window)

	_, err := pipe.Exec(rl.context)
	if err != nil {
		// métrica para saber que foi erro do redis
		return false
	}

	return incr.Val() <= int64(rl.limit)
}

func rateLimiterMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		if !rl.Allow(clientIP) {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
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

	handler := rateLimiterMiddleware(rateLimiter, router)

	http.ListenAndServe(":8080", handler)
}

// docker run -d --name redislimiter -p 6379:6379 redis:latest
