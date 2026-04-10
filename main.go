package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// 👉 Função separada (testável)
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

// 👉 Função separada (testável)
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello ARGOCD + Jaeger + OpenTelemetry 🚀")
}

func initTracer() func(context.Context) error {
	ctx := context.Background()

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("simple-collector.observability.svc:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("erro ao criar exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("goapp"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func main() {
	shutdown := initTracer()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = shutdown(ctx)
	}()

	mux := http.NewServeMux()

	// 👉 usando funções separadas
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/ping", PingHandler)

	handler := otelhttp.NewHandler(mux, "goapp-http")

	log.Println("🚀 server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
