package gengql

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/gengql/generated"
	"github.com/99designs/gqlgen/graphql"
	"go.opentelemetry.io/otel"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var tracer = otel.Tracer("gengql")

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
