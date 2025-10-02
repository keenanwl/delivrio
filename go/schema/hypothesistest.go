package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HypothesisTest struct {
	ent.Schema
}

func (HypothesisTest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (HypothesisTest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hypothesis_test_delivery_option", HypothesisTestDeliveryOption.Type).
			Unique(),
		edge.To("connection", Connection.Type).
			Unique().
			Required(),
	}
}

func (HypothesisTest) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (HypothesisTest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("hypothesis_test")),
	}
}

func (HypothesisTest) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("active").
			Default(false),
	}
}
