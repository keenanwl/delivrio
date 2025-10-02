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

type CarrierDF struct {
	ent.Schema
}

func (CarrierDF) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierDF) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_df").
			Unique().
			Required(),
	}
}

func (CarrierDF) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierDF) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_df")),
	}
}

func (CarrierDF) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_id"),
		field.String("agreement_number"),
		field.Enum("who_pays").
			Values("Prepaid", "Collect").
			Default("Prepaid"),
		field.Bool("test").
			Default(true),
	}
}
