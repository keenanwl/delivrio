package limiter

import (
	"context"

	"delivrio.io/go/ent/intercept"
	"entgo.io/ent"
)

func Limiter(limit int) ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			q, err := intercept.NewQuery(query)
			if err != nil {
				return nil, err
			}
			q.Limit(limit)
			return next.Query(ctx, query)
		})
	})
}
