package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_requests",
			Help: "Total number of requests",
		},
		[]string{"endpoint"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration",
			Help:    "Duração da requisição",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
)

func init() {
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(requestDuration)
}

func main() {

	tp, err := initTracing()
	if err != nil {
		log.Fatalf("could not initialize tracing: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("could not shutdown tracing: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("/process", otelhttp.NewHandler(http.HandlerFunc(processHandler), "processHandler"))
	mux.Handle("/metrics", promhttp.Handler())

	port := 8081
	log.Printf("Listening on http://localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	totalRequests.WithLabelValues(r.URL.Path).Inc()
	requestDuration.WithLabelValues("/process").Observe(time.Since(start).Seconds())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func initTracing() (*trace.TracerProvider, error) {

	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, fmt.Errorf("could not create jaeger exporter: %w", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ExampleService"),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

// commit of the day
