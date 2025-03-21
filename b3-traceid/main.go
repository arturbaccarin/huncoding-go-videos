// https://youtu.be/T1CXU-F_VmM
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

var tracer = otel.Tracer("B3-Tracer")

func initTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("my-service"),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(b3.New()))

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cleanup := initTracer()
	defer cleanup()

	go func() {
		http.HandleFunc("/trace", func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(
				r.Context(), propagation.HeaderCarrier(r.Header))

			ctx, span := tracer.Start(ctx, "ReceiveHandler")
			defer span.End()

			tracerId := span.SpanContext().TraceID().String()
			fmt.Printf("TraceID gerado/recebido: %s", tracerId)

			w.Header().Set("x-b3-traceid", tracerId)
			w.Header().Set("b3-traceid", tracerId)
		})

		fmt.Println("Servidor iniciado na porta 8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
