package tracer

import (
	"context"
	"io"
	"time"
	"weather/pkg/config"
	"weather/pkg/logger"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

type Tracer interface {
	StartSpanFromContext(ctx context.Context, name string, ops ...opentracing.StartSpanOption) (opentracing.Span, context.Context)
}

type tracer struct {
	t opentracing.Tracer
}

var closer io.Closer

func NewTracer(configer config.Configer, logger logger.Logger) Tracer {
	cfg := jaegercfg.Configuration{
		ServiceName: configer.GetString("SERVICE_NAME"),
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			// get from config
			// LocalAgentHostPort: ,
		},
	}

	t, c, err := cfg.NewTracer(
		jaegercfg.Logger(jaeger.StdLogger),
	)

	if err != nil {
		panic(err)
	}

	closer = c
	opentracing.SetGlobalTracer(t)
	return &tracer{
		t: t,
	}
}

func Close() error {
	return closer.Close()
}

func (t *tracer) StartSpanFromContext(ctx context.Context, name string, ops ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, name, ops...)
}
