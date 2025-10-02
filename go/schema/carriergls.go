package schema

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/carriergls"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type CarrierGLS struct {
	ent.Schema
}

func (CarrierGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

type ErrCarrierGLSFields struct {
	invalidFields map[string]string
}

func (ErrCarrierGLSFields) Error() string {
	return "one or more CarrierGLS fields is invalid"
}
func (e ErrCarrierGLSFields) String() string {
	return e.Error()
}
func (e ErrCarrierGLSFields) InvalidFields(ctx context.Context) []*gqlerror.Error {
	var out = make([]*gqlerror.Error, 0)
	for f, e := range e.invalidFields {
		out = append(out, &gqlerror.Error{
			Path:    graphql.NewPathWithField(f).Path(),
			Message: e,
		})
	}
	return out
}

func NewErrCarrierGLSFields() ErrCarrierGLSFields {
	return ErrCarrierGLSFields{
		invalidFields: map[string]string{},
	}
}

func (CarrierGLS) Hooks() []ent.Hook {

	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CarrierGLSFunc(func(ctx context.Context, m *ent2.CarrierGLSMutation) (ent.Value, error) {
					err := NewErrCarrierGLSFields()
					if num, ok := m.GLSPassword(); ok && len(num) < 5 {
						err.invalidFields[camel(carriergls.FieldGLSPassword)] = "password must be greater than 5 characters"
					}
					if num, ok := m.GLSUsername(); ok && len(num) < 5 {
						err.invalidFields[camel(carriergls.FieldGLSUsername)] = "GLS username must be greater than 5 characters"
					}

					if len(err.invalidFields) == 0 {
						return next.Mutate(ctx, m)
					}
					return nil, err
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}

}

func (CarrierGLS) Fields() []ent.Field {
	return []ent.Field{
		field.String("contact_id").
			Optional(),
		field.String("gls_username").
			Optional(),
		field.String("gls_password").
			Optional(),
		field.String("customer_id").
			Optional(),
		field.String("gls_country_code").
			Optional(),
		field.Bool("sync_shipment_cancellation").
			Optional().
			Default(false),
		field.Bool("print_error_on_label").
			Optional().
			Default(false),
	}
}

func (CarrierGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_gls").
			Unique().
			Required(),
	}
}

func (CarrierGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_gls")),
	}
}
