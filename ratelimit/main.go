package main

import (
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	return ip
}

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 3) // 3 requisições por segundo
		visitors[ip] = limiter
	}

	return limiter
}

func rateLimitByIP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		limiter := getVisitor(ip)
		if !limiter.Allow() {
			log.Printf("[BLOCKED] IP: %s\n", ip)
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		log.Printf("[ALLOWED] IP: %s\n", ip)
		h.ServeHTTP(w, r)
	})
}

func doRequestUsingDifferentIPs() {
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
	wg := sync.WaitGroup{}

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
				req.RemoteAddr = ip + ":12345"
				_, _ = http.DefaultClient.Do(req)
			}
		}(ip)
	}

	wg.Wait()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	go func() {
		http.ListenAndServe(":8080", rateLimitByIP(mux))
	}()

	doRequestUsingDifferentIPs()
}

// func main() {
// 	http.Handle("/", rateLimitByIP(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello World"))
// 	})))

// 	http.ListenAndServe(":8080", nil)
// }
