package logger

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

const (
	keyTraceID = "trace_id"
	keySpanID  = "span_id"
)

var logger = NewLogger(WithDebugLevel(), WithOutputStdout())

func NewLogger(opts ...ConfigOption) *Logger {
	logger, err := NewConfig().Build()
	if err != nil {
		panic(err)
	}

	return &Logger{
		logger: logger,
	}
}

func Close() error {
	if logger != nil && logger.logger != nil {
		err := logger.logger.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Logger) Debug(ctx context.Context, m string) {
	l.logger.Debug(m, getFields(ctx)...)
}

func (l *Logger) Debugf(ctx context.Context, m string, args ...any) {
	l.logger.Debug(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Info(ctx context.Context, m string) {
	l.logger.Info(m, getFields(ctx)...)
}

func (l *Logger) Infof(ctx context.Context, m string, args ...any) {
	l.logger.Info(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Warn(ctx context.Context, m string) {
	l.logger.Warn(m, getFields(ctx)...)
}

func (l *Logger) Warnf(ctx context.Context, m string, args ...any) {
	l.logger.Warn(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Error(ctx context.Context, m string) {
	l.logger.Error(m, getFields(ctx)...)
}

func (l *Logger) Errorf(ctx context.Context, m string, args ...any) {
	l.logger.Error(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Panic(ctx context.Context, m string) {
	l.logger.Panic(m, getFields(ctx)...)
}

func (l *Logger) Panicf(ctx context.Context, m string, args ...any) {
	l.logger.Panic(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func (l *Logger) Fatal(ctx context.Context, m string) {
	l.logger.Fatal(m, getFields(ctx)...)
}

func (l *Logger) Fatalf(ctx context.Context, m string, args ...any) {
	l.logger.Fatal(fmt.Sprintf(m, args...), getFields(ctx)...)
}

func Debug(ctx context.Context, m string) {
	logger.Debug(ctx, m)
}

func Debugf(ctx context.Context, m string, args ...any) {
	logger.Debugf(ctx, m, args...)
}

func Info(ctx context.Context, m string) {
	logger.Info(ctx, m)
}

func Infof(ctx context.Context, m string, args ...any) {
	logger.Infof(ctx, m, args...)
}

func Warn(ctx context.Context, m string) {
	logger.Info(ctx, m)
}

func Warnf(ctx context.Context, m string, args ...any) {
	logger.Warnf(ctx, m, args...)
}

func Error(ctx context.Context, m string) {
	logger.Error(ctx, m)
}

func Errorf(ctx context.Context, m string, args ...any) {
	logger.Errorf(ctx, m, args...)
}

func Panic(ctx context.Context, m string) {
	logger.Panic(ctx, m)
}

func Panicf(ctx context.Context, m string, args ...any) {
	logger.Panicf(ctx, m, args...)
}

func Fatal(ctx context.Context, m string) {
	logger.Fatal(ctx, m)
}

func Fatalf(ctx context.Context, m string, args ...any) {
	logger.Fatalf(ctx, m, args...)
}

func getFields(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0, 2)
	traceID := getTraceID(ctx)
	if traceID != "" {
		fields = append(fields, zap.String(keyTraceID, traceID))
	}
	spanID := getSpanID(ctx)
	if spanID != "" {
		fields = append(fields, zap.String(keySpanID, spanID))
	}
	return fields
}

func getTraceID(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}

func getSpanID(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}
	return ""
}
