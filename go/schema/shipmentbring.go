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

type ShipmentBring struct {
	ent.Schema
}

func (ShipmentBring) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentBring) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_bring").
			Unique().
			Required(),
	}
}

func (ShipmentBring) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentBring) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_bring")),
	}
}

func (ShipmentBring) Fields() []ent.Field {
	return []ent.Field{
		field.String("consignment_number"),
	}
}
