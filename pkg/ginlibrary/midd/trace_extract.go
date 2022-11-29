package midd

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/otel-trace-demo/pkg/idgen"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

var (
	gener idgen.IDGenerator
)

func init() {
	gener = idgen.NewDefaultGenerator()
}

const (
	ginOtelTraceSpan     = `ginOpenTelemetrySpan`
	ginOtelTraceParentID = `ginOtelTraceParentID`

	ginLogger = `ginLogger`
)

// TraceSpanExtractMiddleware 一个 middleware。
// 解析 header, 并在 gin value 中加入 OpenTelemetry trace span
func TraceSpanExtractMiddleware(c *gin.Context) {

	// 1. 从 header 中获取关键字
	mapHeader := propagation.MapCarrier{
		"traceparent": c.GetHeader("traceparent"),
		"tracestate":  c.GetHeader("tracestate"),
	}

	// 2.1. 定义 propagator
	pp := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	)

	// 2.2. 从 header 中提取 span
	ctx := pp.Extract(c, mapHeader)
	span := trace.SpanFromContext(ctx)
	spanCtx := span.SpanContext()

	// 2.3. 在 gin context value 中设置 parent_id
	c.Set(ginOtelTraceParentID, spanCtx.SpanID().String())

	// 2.4. 根据 span 获取新 span
	span = newSpan(span)

	// 2.5. 将 span 加入到 gin context value 中
	c.Set(ginOtelTraceSpan, span)
	c.Next()
}

// TraceSpanFromContext 从 context 中获取 trace span
// 这里需要封装一下
//
// trace.SpanFromContext 使用的是私有key:
// .   const currentSpanKey traceContextKeyType = iota
// 在 gin 中只能使用 string 作为 key， 我们自己也创建了一个私有key
// .   const ginOtelTraceSpan = `ginOpenTelemetrySpan`
//
// 为了能避免 key 造成的影响， 因此需要对 gin.Context 做一个判断。
func TraceSpanFromContext(ctx context.Context) trace.Span {

	// 判断是否为 gin.Context
	if ginctx, ok := ctx.(*gin.Context); ok {
		v := ginctx.Value(ginOtelTraceSpan)
		if span, ok := v.(trace.Span); ok {
			return span
		}
	}

	// 使用 trace 的原始封装
	return trace.SpanFromContext(ctx)
}

// newSpan 根据 TraceID 是否合法， 返回一个新 span
// 如果 TraceId 合法， 则变更 SpanID
// 如果 TraceID 不合法， 则创建一个全新 Span
func newSpan(span trace.Span) trace.Span {
	spanCtx := span.SpanContext()
	ctx := context.Background()

	if spanCtx.TraceID().IsValid() {
		sid := gener.NewSpanID(ctx, span.SpanContext().TraceID())
		spanCtx = spanCtx.WithSpanID(sid)
		ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)
	} else {
		tid, sid := gener.NewIDs(ctx)
		spanCtx = spanCtx.WithSpanID(sid)
		spanCtx = spanCtx.WithTraceID(tid)
		ctx = trace.ContextWithSpanContext(ctx, spanCtx)
	}

	return trace.SpanFromContext(ctx)
}

// traceParentIDFromContext 从 从 gin context 中获取 trace parend id
func traceParentIDFromContext(ctx context.Context) string {
	v := ctx.Value(ginOtelTraceParentID)

	id, ok := v.(string)
	if ok {
		return id
	}
	return ""
}
