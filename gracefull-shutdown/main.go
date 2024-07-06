package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		time.Sleep(4 * time.Second)
	})

	// router.Run("8080")

	chanError := make(chan error)

	go GracefullyShutdown(router, "8080", chanError)

	if err := <-chanError; err != nil {
		fmt.Println(err)
	}

}

func GracefullyShutdown(handler http.Handler, port string, chanError chan error) {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		fmt.Println("received a shutdown signal, quiting...")

		shutdownTimeout := 5 * time.Second
		ctxTimeout, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

		defer func() {
			stop()
			cancel()
			close(chanError)
		}()

		err := server.Shutdown(ctxTimeout)
		if err != nil {
			fmt.Println("shutdown completed")
		}
	}()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			chanError <- fmt.Errorf("an error occurre whil trying to start application: %w", err)
			return
		}
	}()
}
