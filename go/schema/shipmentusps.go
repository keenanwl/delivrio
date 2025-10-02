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

type ShipmentUSPS struct {
	ent.Schema
}

func (ShipmentUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_usps").
			Unique().
			Required(),
	}
}

func (ShipmentUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_usps")),
	}
}

func (ShipmentUSPS) Fields() []ent.Field {
	return []ent.Field{
		field.String("tracking_number").
			Optional(),
		field.Float("postage").
			Optional(),
		field.Time("scheduled_delivery_date").
			Optional(),
	}
}
