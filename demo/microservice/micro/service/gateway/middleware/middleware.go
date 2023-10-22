package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

func JaegerGatewayMiddleware(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var md = make(metadata.Metadata, 1)
		opName := ctx.Request.URL.Path + "-" + ctx.Request.Method // 操作名称
		parentSpan := tracer.StartSpan(opName)
		defer parentSpan.Finish()
		injectErr := tracer.Inject(parentSpan.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)) // 将TraceID注入到md中
		if injectErr != nil {
			logger.Fatalf("%s: Couldn't inject metadata", injectErr)
		}
		newCtx := metadata.NewContext(ctx.Request.Context(), md) // 利用context传递TraceID
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
