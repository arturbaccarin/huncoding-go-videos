// https://youtu.be/4vYuZrcTInQ 3:45

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/pprof"
	"syscall"
	"time"
)

func main() {
	cpuProfile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("could not create CPU profile: %v", err)
	}

	if err := pprof.StartCPUProfile(cpuProfile); err != nil {
		log.Fatalf("could not start CPU profile: %v", err)
	}

	// Run the program
	http.HandleFunc("/compute", computeHandler)
	http.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	stopChan := make(chan os.Signal, 1)
	signalNotify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Listening on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on server: %v", err)
		}
	}()

	<-stopChan
	log.Println("Signal received, shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown server: %v", err)
	}

	pprof.StopCPUProfile()
	cpuProfile.Close()
	log.Println("CPU profile saved to cpu.prof")

	log.Println("Goodbye!")
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	n := rand.Intn(1000)
	result := fib(n % 30)

	duration := time.Since(start)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"input": %d, "result": %d, "duration": %d}`, n, result, duration.Milliseconds())))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
