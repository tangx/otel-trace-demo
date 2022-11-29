package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tangx/otel-trace-demo/pkg/ginlibrary/midd"
	"github.com/tangx/otel-trace-demo/pkg/httpclient"
)

func main() {
	r := gin.Default()

	r.Use(
		midd.TraceSpanExtractMiddleware,
		midd.LoggerWithTraceSpanMiddleware,
	)
	r.Use(midd.TraceSpanInjectMiddleware)

	r.GET("/", pingpong)

	err := r.Run(":8088")

	if err != nil {
		panic(err)
	}

}

func pingpong(c *gin.Context) {
	log := midd.LoggerFromContext(c)
	span := midd.TraceSpanFromContext(c)

	log.Info("SERVER1", "kk", "vv")

	b, err := span.SpanContext().MarshalJSON()
	if err != nil {
		c.String(500, "failed span")
		return
	}

	// ctx := trace.ContextWithSpan(c, span)
	_ = reqServer2(c)

	c.String(200, string(b))
}

func reqServer2(ctx context.Context) error {
	return httpclient.GET(ctx, `http://127.0.0.1:9099/`)
}
