package logger

import (
	"context"
	"weather/pkg/logger/utils"

	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap/zapcore"
)

// DebugWithContext logs on debug level and trace based on the context span if it exists
func (l *logger) DebugWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.DebugWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// DebugWithSpan logs on debug level and add the logs on the trace if span exists.
func (l *logger) DebugWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.Debug(log, fields...)
	l.logSpan(span, log, fields...)
}

// Debug logs on debug level
func (l *logger) Debug(log string, fields ...zapcore.Field) {
	l.l.Debug(log, fields...)
}

// INFO

// InfoWithContext logs on info level and trace based on the context span if it exists
func (l *logger) InfoWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.InfoWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// InfoWithSpan logs on info level and add the logs on the trace if span exists.
func (l *logger) InfoWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.Info(log, fields...)
	l.logSpan(span, log, fields...)

}

// Info logs on info level
func (l *logger) Info(log string, fields ...zapcore.Field) {
	l.l.Info(log, fields...)
}

// WARN

// WarnWithContext logs on warn level and trace based on the context span if it exists
func (l *logger) WarnWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.WarnWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// WarnWithSpan logs on warn level and add the logs on the trace if span exists.
func (l *logger) WarnWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.Warn(log, fields...)
	l.logSpan(span, log, fields...)

}

// Warn logs on warn level
func (l *logger) Warn(log string, fields ...zapcore.Field) {
	l.l.Warn(log, fields...)
}

// ERROR

// ErrorWithContext logs on error level and trace based on the context span if it exists
func (l *logger) ErrorWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.ErrorWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// ErrorWithSpan logs on error level and add the logs on the trace if span exists.
func (l *logger) ErrorWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.Error(log, fields...)
	l.logSpan(span, log, fields...)
}

// Error logs on error level
func (l *logger) Error(log string, fields ...zapcore.Field) {
	l.l.Error(log, fields...)
}

// DPANIC

// DPanicWithContext logs on dPanic level and trace based on the context span if it exists
func (l *logger) DPanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.DPanicWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// DPanicWithSpan logs on dPanic level and add the logs on the trace if span exists.
func (l *logger) DPanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.logSpan(span, log, fields...)
	l.DPanic(log, fields...)
}

// DPanic logs on dPanic level
func (l *logger) DPanic(log string, fields ...zapcore.Field) {
	l.l.DPanic(log, fields...)
}

// PANIC

// PanicWithContext logs on panic level and trace based on the context span if it exists
func (l *logger) PanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.PanicWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// PanicWithSpan logs on panic level and add the logs on the trace if span exists.
func (l *logger) PanicWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.logSpan(span, log, fields...)
	l.Panic(log, fields...)
}

// Panic logs on panic level
func (l *logger) Panic(log string, fields ...zapcore.Field) {
	l.l.Panic(log, fields...)
}

// FatalWithContext logs on fatal level and trace based on the context span if it exists
func (l *logger) FatalWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	l.FatalWithSpan(opentracing.SpanFromContext(ctx), log, fields...)
}

// FatalWithSpan logs on fatal level and add the logs on the trace if span exists.
func (l *logger) FatalWithSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	l.logSpan(span, log, fields...)
	l.Fatal(log, fields...)
}

// Fatal logs on fatal level
func (l *logger) Fatal(log string, fields ...zapcore.Field) {
	l.l.Fatal(log, fields...)
}

func (l *logger) logSpan(span opentracing.Span, log string, fields ...zapcore.Field) {
	if span != nil {
		opentracingFields := make([]opentracinglog.Field, len(fields)+1)
		if log != "" {
			opentracingFields = append(opentracingFields, opentracinglog.String("event", log))
		}
		if len(fields) > 0 {
			opentracingFields = append(opentracingFields, utils.ZapFieldsToOpentracing(fields...)...)
		}
		span.LogFields(opentracingFields...)
	}
}
