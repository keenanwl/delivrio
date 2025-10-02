package hooks

import (
	"context"

	"entgo.io/ent/entc/gen"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	camel = gen.Funcs["camel"].(func(string) string)
)

type ValidationError struct {
	basePath      string
	invalidFields map[string]string
}

func (e ValidationError) SetBasePath(path string) {
	e.basePath = path
}
func (e ValidationError) SetError(path, err string) {
	e.invalidFields[camel(path)] = err
}
func (ValidationError) Error() string {
	return "one or more fields are invalid"
}
func (e ValidationError) String() string {
	return e.Error()
}
func (e ValidationError) InvalidFields(ctx context.Context) []*gqlerror.Error {
	var out = make([]*gqlerror.Error, 0)
	for f, e := range e.invalidFields {
		out = append(out, &gqlerror.Error{
			Path:    graphql.GetPathContext(graphql.WithPathContext(ctx, graphql.NewPathWithField(f))).Path(),
			Message: e,
		})
	}
	return out
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		invalidFields: map[string]string{},
	}
}
