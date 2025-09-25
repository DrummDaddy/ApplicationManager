package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracing(serviceName string) func() {
	exp, _ := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
	otel.SetTracerProvider(tracerProvider)

	return func() {
		_ = tracerProvider.Shutdown(context.Background())
	}

}
