package tracer

import (
	"context"
	"fmt"
	"tinvest-go/internal/config"

	"go.opentelemetry.io/otel"
	//nolint:staticcheck
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk_trace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	otel_trace "go.opentelemetry.io/otel/trace"
)

var tracer otel_trace.Tracer = NewTracer()

func NewTracer() otel_trace.Tracer {
	exporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint(config.JaegerUrl),
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to create jaeger exporter: %w", err))
	}

	tracerProvider := sdk_trace.NewTracerProvider(
		sdk_trace.WithBatcher(exporter),
		sdk_trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.AppName),
		)),
	)

	otel.SetTracerProvider(tracerProvider)

	tracer := otel.Tracer("default")

	return tracer
}

func StartSpanFromContext(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, otel_trace.Span) {
	ctx, span := tracer.Start(ctx, name, opts...)
	return ctx, span
}

func Close() error {
	var tracerProvider *sdk_trace.TracerProvider
	tracerProvider, _ = otel.GetTracerProvider().(*sdk_trace.TracerProvider)
	err := tracerProvider.ForceFlush(context.Background())
	if err != nil {
		return err
	}
	err = tracerProvider.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return nil
}
