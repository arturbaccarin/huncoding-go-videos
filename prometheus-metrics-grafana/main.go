package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// this is the commit of the day :)
var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total", // o nome da métrica
			Help: "number of requests",
		}, []string{"path"}) // dimensões da métrica

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "duration of http request",
		}, []string{"path"})
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
}

func handler(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(requestDuration.WithLabelValues(r.URL.Path))
	defer timer.ObserveDuration()

	httpRequests.WithLabelValues(r.URL.Path)

	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", nil)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Application up and running")
	http.ListenAndServe(":8080", nil)
}
