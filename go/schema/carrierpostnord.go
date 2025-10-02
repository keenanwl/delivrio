package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CarrierPostNord struct {
	ent.Schema
}

func (CarrierPostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierPostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_post_nord").
			Unique().
			Required(),
	}
}

func (CarrierPostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierPostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_post_nord")),
	}
}

func (CarrierPostNord) Fields() []ent.Field {
	return []ent.Field{
		field.String("customer_number").
			Default("").
			Comment("Default empty to allow creation from dialog with followup editing"),
	}
}
