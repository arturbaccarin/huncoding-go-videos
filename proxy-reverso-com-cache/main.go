// https://youtu.be/_BPXJaEEjKc (5min)
package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Cache struct {
	data map[string][]byte
	ttl  map[string]time.Time
	mu   sync.RWMutex
}

type ReverseProxy struct {
	routes map[string][]string
	cache  *Cache
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
		ttl:  make(map[string]time.Time),
	}
}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{
		routes: map[string][]string{
			"/todos/1": {
				"http://jsonplaceholder.typicode.com",
				"http://jsonplaceholder.typicode.com",
			},
		},
		cache: NewCache(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if expiration, exist := c.ttl[key]; exist && time.Now().Before(expiration) {
		return c.data[key], true
	}

	return nil, false
}

func (c *Cache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	c.ttl[key] = time.Now().Add(ttl)
}

func (c *Cache) CleanUp() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, expiration := range c.ttl {
		if time.Now().After(expiration) {
			delete(c.data, key)
			delete(c.ttl, key)
		}
	}
}

func (rp *ReverseProxy) cacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := fmt.Sprintf("%s-%x", r.URL.Path, sha256.Sum256([]byte(r.URL.RawQuery)))

		if cache, ok := rp.cache.Get(key); ok {
			w.Write(cache)
			fmt.Printf("Cache hit: %s\n", r.URL.Path)
			return
		}

		recorder := &responseRecorder{
			ResponseWriter: w,
			body:           bytes.NewBuffer(nil),
		}
		next(recorder, r)

		rp.cache.Set(key, recorder.body.Bytes(), 5*time.Second)
	}
}

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (rp *ReverseProxy) selectBackend(path string) (string, bool) {
	backend, exists := rp.routes[path]
	if !exists || len(backend) == 0 {
		return "", false
	}

	return backend[rand.Intn(len(backend))], true
}

func transformBody(body []byte) []byte {
	return bytes.ReplaceAll(body, []byte("userId"), []byte("user_id"))
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, exists := rp.selectBackend(r.URL.Path)
	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	remote, err := url.Parse(backend)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	proxyReq, err := http.NewRequest(r.Method, remote.String()+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	proxyReq.Header = r.Header

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	body := transformBody(responseBytes)

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	log.Printf("Request: %s, Backend: %s, Status: %d", r.URL.Path, backend, resp.StatusCode)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	proxy := NewReverseProxy()

	rp := NewReverseProxy()
	http.HandleFunc("/", proxy.cacheMiddleware(rp.ServeHTTP))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
