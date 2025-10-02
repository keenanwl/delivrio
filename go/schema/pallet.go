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

type Pallet struct {
	ent.Schema
}

func (Pallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Pallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type),
		edge.To("packaging", Packaging.Type).
			Unique(),
		edge.From("consolidation", Consolidation.Type).
			Ref("pallets").
			Unique().
			Required(),
		edge.To("shipment_pallet", ShipmentPallet.Type).
			Unique().
			Comment("A pallet may only have 1 active shipment"),
		edge.To("cancelled_shipment_pallet", ShipmentPallet.Type).
			Comment("Cancelled shipments move here to maintain the ref"),
	}
}

func (Pallet) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Pallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("pallet")),
	}
}

func (Pallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("public_id"),
		field.String("description"),
	}
}
