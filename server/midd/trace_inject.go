package midd

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

func TraceSpanInjectMiddleware(c *gin.Context) {

	// 1. 从 gin context value 中获取 span
	span := TraceSpanFromContext(c)
	if !span.SpanContext().IsValid() {
		// 如果 span 不合法， 直接跳过
		c.Next()
	}

	// 2. 创建输入容器和处理器
	headerCarrier := propagation.MapCarrier{}
	pp := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	)

	// 3. 创建可以被 trace 处理的 ctx
	ctx := trace.ContextWithSpan(context.Background(), span)

	// 4. 注入数据
	pp.Inject(ctx, headerCarrier)

	// 5. 写入 response header
	for k, v := range headerCarrier {
		c.Writer.Header().Set(k, v)
	}

	c.Next()
}
