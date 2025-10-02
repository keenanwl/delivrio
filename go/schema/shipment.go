package schema

import (
	"delivrio.io/go/schema/hooks/shipmenthooks"
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Shipment struct {
	ent.Schema
}

func (Shipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Shipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("shipment_public_id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		field.Enum("status").Values(
			"Pending",
			"Prebooked",
			"Booked",
			"Partially_dispatched",
			"Dispatched",
			"Deleted",
		),
	}
}

func (Shipment) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Shipment) Hooks() []ent.Hook {
	return []ent.Hook{
		shipmenthooks.UpdateShipmentConnectedEntities(),
	}
}

func (Shipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier", Carrier.Type).
			Comment("Can get this from edges, but want the reference to persist even after cancelling").
			Unique().
			Required().
			Immutable(),
		edge.To("shipment_history", ShipmentHistory.Type),
		edge.To("shipment_bring", ShipmentBring.Type).
			Unique(),
		edge.To("shipment_dao", ShipmentDAO.Type).
			Unique(),
		edge.To("shipment_df", ShipmentDF.Type).
			Unique(),
		edge.To("shipment_dsv", ShipmentDSV.Type).
			Unique(),
		edge.To("shipment_easy_post", ShipmentEasyPost.Type).
			Unique(),
		edge.To("shipment_post_nord", ShipmentPostNord.Type).
			Unique(),
		edge.To("shipment_gls", ShipmentGLS.Type).
			Unique(),
		edge.To("shipment_usps", ShipmentUSPS.Type).
			Unique(),

		edge.To("consolidation", Consolidation.Type).
			Unique().
			Comment("A shipment may have 0 or more collis"),

		edge.To("old_consolidation", Consolidation.Type).
			Comment("After a shipment is cancelled"),

		edge.To("shipment_parcel", ShipmentParcel.Type).
			Comment("A shipment may have 0 or more collis"),

		edge.To("shipment_pallet", ShipmentPallet.Type).
			Comment("A shipment may have 0 or more pallets"),
	}
}

func (Shipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment")),
	}
}
