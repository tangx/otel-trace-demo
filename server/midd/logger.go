package midd

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// LoggerWithTraceSpanMiddleware 一个 gin middlware， 在 gin value 中假如预设 logger
func LoggerWithTraceSpanMiddleware(c *gin.Context) {
	log := slog.Default()

	span := TraceSpanFromContext(c)

	log = log.With(
		"parent_id", traceParentIDFromContext(c),
		"span_id", span.SpanContext().SpanID().String(),
		"trace_id", span.SpanContext().TraceID().String(),
	)

	c.Set(ginLogger, log)
	c.Next()
}

// LoggerFromContext 从 gin.Context 中获取 slog.Logger
func LoggerFromContext(c *gin.Context) *slog.Logger {
	log, ok := c.Value(ginLogger).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return log
}
