package zipkin

import (
	"github.com/opentracing/opentracing-go"
)

func NewZipkinTracer(serviceName string, host string) opentracing.Tracer {
	tracer := zipkin.GetTracer(serviceName, host)
	opentracing.SetGlobalTracer(tracer)

	return tracer
}
