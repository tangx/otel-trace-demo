package httpclient

import (
	"context"
	"net/http"

	"github.com/tangx/otel-trace-demo/pkg/ginlibrary/midd"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func GET(ctx context.Context, ur string) error {

	headers := TraceMapCarrier(ctx)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ur, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.DefaultClient
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func TraceMapCarrier(ctx context.Context) propagation.MapCarrier {
	span := midd.TraceSpanFromContext(ctx)
	ctx = trace.ContextWithSpan(ctx, span)

	pp := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	)

	carrier := propagation.MapCarrier{}
	pp.Inject(ctx, carrier)

	return carrier
}
