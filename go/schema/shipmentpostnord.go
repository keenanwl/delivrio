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

type ShipmentPostNord struct {
	ent.Schema
}

func (ShipmentPostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentPostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_post_nord").
			Unique().
			Required(),
	}
}

func (ShipmentPostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentPostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_post_nord")),
	}
}

func (ShipmentPostNord) Fields() []ent.Field {
	return []ent.Field{
		field.String("booking_id").
			Comment("Multiple labels. May contain unrelated shipments."),
		field.String("item_id").
			Comment("Individual label. Can be grouped to same address. Probably should not be on this ent?"),
		field.String("shipment_reference_no").
			Comment("Shipment can contain multiple parcels to same address."),
	}
}
