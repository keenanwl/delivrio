package schema

import (
	"delivrio.io/go/schema/hooks/hookshipmentparcel"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ShipmentParcel struct {
	ent.Schema
}

func (ShipmentParcel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ShipmentParcel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shipment", Shipment.Type).
			Ref("shipment_parcel").
			Unique().
			Required().
			Annotations(
				entgql.OrderField("SHIPMENT_CREATED_AT"),
			).
			Immutable(),
		edge.From("colli", Colli.Type).
			Ref("shipment_parcel").
			Unique(),
		edge.From("old_colli", Colli.Type).
			Ref("cancelled_shipment_parcel").
			Comment("After shipment cancelled, ref moved here."),
		edge.From("workspace_recent_scan", WorkspaceRecentScan.Type).
			Ref("shipment_parcel"),
		edge.To("packaging", Packaging.Type).
			Unique(),
		edge.From("print_job", PrintJob.Type).
			Ref("shipment_parcel"),
		edge.To("document_file", DocumentFile.Type).
			Unique(),
	}
}

func (ShipmentParcel) Hooks() []ent.Hook {
	return []ent.Hook{
		hookshipmentparcel.CreateShipmentParcelDeliveryEstimate(),
		hookshipmentparcel.CreateShipmentChangeColliStatus(),
		hookshipmentparcel.UpdateShipmentStatusOnShipmentParcelMutation(),
	}
}

func (ShipmentParcel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("shipment_parcel")),
	}
}

func (ShipmentParcel) Fields() []ent.Field {
	return []ent.Field{
		field.String("item_id").
			Optional(),
		field.Enum("status").
			Default("pending").
			Values("pending", "printed", "in_transit", "out_for_delivery", "delivered", "awaiting_cc_pickup", "picked_up"),
		field.Strings("cc_pickup_signature_urls").
			Optional(),
		field.Time("expected_at").
			Optional().
			Annotations(
				entgql.OrderField("EXPECTED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
		field.Time("fulfillment_synced_at").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Time("cancel_synced_at").
			Optional().
			Nillable().
			Comment("For supported carriers will attempt to cancel shipment via the API").
			Annotations(
				entgql.OrderField("CANCEL_SYNCED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}
