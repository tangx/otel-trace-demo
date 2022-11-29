package httpclient

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/propagation"
)

func GET(ctx context.Context, ur string) error {

	headers := XXX(ctx)

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

func XXX(ctx context.Context) propagation.MapCarrier {
	pp := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	)

	carrier := propagation.MapCarrier{}
	pp.Inject(ctx, carrier)

	return carrier
}
