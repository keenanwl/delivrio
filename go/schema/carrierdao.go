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

type CarrierDAO struct {
	ent.Schema
}

func (CarrierDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_dao").
			Unique().
			Required(),
	}
}

func (CarrierDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_dao")),
	}
}

func (CarrierDAO) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_id").
			Optional(),
		field.String("api_key").
			Optional(),
		field.Bool("Test").
			Default(true),
	}
}
