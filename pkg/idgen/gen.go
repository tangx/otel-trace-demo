package idgen

// var gen IDGenerator

// func init() {
// 	gen = defaultIDGenerator()
// }

// func NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
// 	return gen.NewIDs(ctx)
// }

// func NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID {
// 	return gen.NewSpanID(ctx, traceID)
// }

// func NewIDsFromString(ctx context.Context, traceId string, spanId string) (trace.TraceID, trace.SpanID) {
// 	traceID, err := trace.TraceIDFromHex(traceId)
// 	if err != nil {
// 		return NewIDs(ctx)
// 	}

// 	spanID, err := trace.SpanIDFromHex(spanId)
// 	if err != nil {
// 		return traceID, NewSpanID(ctx, traceID)
// 	}

// 	return traceID, spanID
// }

func NewDefaultGenerator() IDGenerator {
	return defaultIDGenerator()
}

// func (gen *randomIDGenerator) NewTraceSpan(ctx context.Context) trace.Span {
// 	traceID, spanID := gen.NewIDs(ctx)

// 	// ss := trace.ContextWithSpanContext(ctx,
// 	//
// 	//	trace.NewSpanContext(trace.SpanContextConfig{}))

// 	span := trace.SpanFromContext(ctx)
// 	span.SpanContext().WithSpanID(spanID)
// 	span.SpanContext().WithTraceID(traceID)

// }
