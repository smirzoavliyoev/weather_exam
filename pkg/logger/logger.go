package logger

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	DebugWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	DebugWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Debug(log string, fields ...zapcore.Field)
	InfoWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	InfoWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Info(log string, fields ...zapcore.Field)
	WarnWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	WarnWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Warn(log string, fields ...zapcore.Field)
	ErrorWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	ErrorWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Error(log string, fields ...zapcore.Field)
	DPanicWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	DPanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	DPanic(log string, fields ...zapcore.Field)
	PanicWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	PanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Panic(log string, fields ...zapcore.Field)
	FatalWithContext(ctx context.Context, log string, fields ...zapcore.Field)
	FatalWithSpan(span opentracing.Span, log string, fields ...zapcore.Field)
	Fatal(log string, fields ...zapcore.Field)
}

type logger struct {
	l *zap.Logger
}

var log logger

// var logger *zap.Logger

func NewLogger() Logger {
	var err error

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"app.log",
	}

	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	log.l = logger
	return &log
}

func Sync() error {
	return log.l.Sync()
}
