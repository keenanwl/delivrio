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

type ShipmentDAO struct {
	ent.Schema
}

func (ShipmentDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_dao").
			Unique().
			Required(),
	}
}

func (ShipmentDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_dao")),
	}
}

func (ShipmentDAO) Fields() []ent.Field {
	return []ent.Field{
		field.String("barcode_id"),
	}
}
