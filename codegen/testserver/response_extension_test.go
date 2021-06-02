package testserver

import (
	"context"
	"testing"

	"github.com/arsmn/fastgql/client"
	"github.com/arsmn/fastgql/graphql"
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/stretchr/testify/require"
)

func TestResponseExtension(t *testing.T) {
	resolvers := &Stub{}
	resolvers.QueryResolver.Valid = func(ctx context.Context) (s string, e error) {
		return "Ok", nil
	}

	srv := handler.NewDefaultServer(
		NewExecutableSchema(Config{Resolvers: resolvers}),
	)

	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		graphql.RegisterExtension(ctx, "example", "value")

		return next(ctx)
	})

	c := client.New(srv.Handler())

	raw, _ := c.RawPost(`query { valid }`)
	require.Equal(t, raw.Extensions["example"], "value")
}
