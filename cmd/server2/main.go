package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/otel-trace-demo/pkg/ginlibrary/midd"
)

func main() {
	r := gin.Default()

	r.Use(
		midd.TraceSpanExtractMiddleware,
		midd.LoggerWithTraceSpanMiddleware,
	)
	r.Use(midd.TraceSpanInjectMiddleware)

	r.GET("/", pingpong)

	err := r.Run(":9099")

	if err != nil {
		panic(err)
	}

}

func pingpong(c *gin.Context) {
	log := midd.LoggerFromContext(c)
	span := midd.TraceSpanFromContext(c)

	log.Info("SERVER2", "kkk", "vvv")

	b, err := span.SpanContext().MarshalJSON()
	if err != nil {
		c.String(500, "failed span")
		return
	}

	c.String(200, string(b))
}
