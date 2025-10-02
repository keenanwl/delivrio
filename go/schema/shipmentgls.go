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

type ShipmentGLS struct {
	ent.Schema
}

func (ShipmentGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ShipmentGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_gls").
			Unique().
			Required(),
	}
}

func (ShipmentGLS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_gls")),
	}
}

func (ShipmentGLS) Fields() []ent.Field {
	return []ent.Field{
		field.String("consignment_id"),
	}
}
