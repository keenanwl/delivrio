package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShipmentHistory holds the schema definition for the ShipmentHistory entity.
type ShipmentHistory struct {
	ent.Schema
}

// Fields of the ShipmentHistory.
func (ShipmentHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Immutable().
			Values("create", "update", "delete"),
	}
}

// Edges of the ShipmentHistory.
func (ShipmentHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_history").
			Unique().
			Required(),
	}
}

func (ShipmentHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ChangeHistoryEntityMixin("shipment"),
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_history")),
	}
}
