package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type ShipmentDF struct {
	ent.Schema
}

func (ShipmentDF) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ShipmentDF) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_df").
			Unique().
			Required(),
	}
}

func (ShipmentDF) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentDF) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_df")),
	}
}

func (ShipmentDF) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
