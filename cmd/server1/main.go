package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/otel-trace-demo/pkg/ginlibrary/midd"
	"github.com/tangx/otel-trace-demo/pkg/httpclient"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{
				"/liveness",
				"/healthy",
			},
		}),
		gin.Recovery(),
	)

	r.Use(
		midd.TraceSpanExtractMiddleware,
		midd.LoggerWithTraceSpanMiddleware,
	)
	r.Use(midd.TraceSpanInjectMiddleware)

	r.GET("/", pingpong)
	r.GET("/liveness", livenessCheck)

	err := r.Run(":80")

	if err != nil {
		panic(err)
	}
}

func livenessCheck(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func pingpong(c *gin.Context) {
	log := midd.LoggerFromContext(c)
	span := midd.TraceSpanFromContext(c)

	log.Info(config.AppName, "kk", "vv")

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

	if config.NextServer == "" {
		return errors.New("NO MORE next server")
	}

	return httpclient.GET(ctx, config.NextServer)
}
