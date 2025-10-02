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

type ShipmentEasyPost struct {
	ent.Schema
}

func (ShipmentEasyPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ShipmentEasyPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_easy_post").
			Unique().
			Required(),
	}
}

func (ShipmentEasyPost) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentEasyPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_easy_post")),
	}
}

func (ShipmentEasyPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("tracking_number").
			Comment("duplicate, may be dropped after verifying").
			Optional(),
		field.String("ep_shipment_id").
			Optional(),
		field.Float("rate").
			Optional(),
		field.Time("est_delivery_date").
			Optional(),
	}
}
