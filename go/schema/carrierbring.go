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

type CarrierBring struct {
	ent.Schema
}

func (CarrierBring) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierBring) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_bring").
			Unique().
			Required(),
	}
}

func (CarrierBring) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierBring) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_bring")),
	}
}

func (CarrierBring) Fields() []ent.Field {
	return []ent.Field{
		field.String("api_key").
			Optional(),
		field.String("customer_number").
			Optional(),
		field.Bool("test").
			Default(true),
	}
}
