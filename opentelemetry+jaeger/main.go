//

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

var tracer = otel.Tracer("Example-Go-Tracer")

func initTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint()) // joga no terminal

	// exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost("localhost"), jaeger.WithAgentPort(6831)))
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("Example-Trce"),
	)

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resources),
	)

	otel.SetTracerProvider(traceProvider)

	return func() {
		err := traceProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cleanup := initTracer()
	defer cleanup()

	ctx, span := tracer.Start(context.Background(), "main")
	defer span.End()

	doWork(ctx)
}

func doWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doWork")
	defer span.End()

	time.Sleep(200 * time.Millisecond)

	doSubWork(ctx)
}

func doSubWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doSubWork")
	defer span.End()

	time.Sleep(500 * time.Millisecond)

	doSubSubWork(ctx)
}

func doSubSubWork(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doSubSubWork")
	defer span.End()

	time.Sleep(1 * time.Second)

	fmt.Println("Done!")
}
