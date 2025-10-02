package apm

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/intercept"
	"fmt"
	"go.opentelemetry.io/otel"
)

func QueryTracer(serverID string) ent.InterceptFunc {
	return func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			q, err := intercept.NewQuery(query)
			if err != nil {
				return nil, err
			}
			traceCtx, span := otel.Tracer(serverID).
				Start(ctx, fmt.Sprintf("ent-query-%v", q.Type()))
			defer span.End()

			value, err := next.Query(traceCtx, query)
			if err != nil {
				span.RecordError(err)
			}

			return value, err
		})
	}
}
