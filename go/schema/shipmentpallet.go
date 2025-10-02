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

type ShipmentPallet struct {
	ent.Schema
}

func (ShipmentPallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentPallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pallet", Pallet.Type).
			Ref("shipment_pallet").
			Unique(),
		edge.From("old_pallet", Pallet.Type).
			Ref("cancelled_shipment_pallet").
			Comment("After shipment cancelled, ref moved here. Mostly for consistency, since the Shipment is also connected still."),
		edge.From("shipment", Shipment.Type).
			Ref("shipment_pallet").
			Unique().
			Required().
			Immutable(),
	}
}

func (ShipmentPallet) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ShipmentPallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_pallet")),
	}
}

func (ShipmentPallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("barcode"),

		// These may be specific to DF and require future refactoring
		field.String("colli_number"),
		field.String("carrier_id"),

		field.String("label_pdf").
			Optional(),
		field.String("label_zpl").
			Optional(),
		field.Enum("status").
			Default("pending").
			Values("pending", "printed"),
	}
}
