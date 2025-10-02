package schema

import (
	"delivrio.io/go/schema/hooks/history"
	"entgo.io/ent/schema/index"
	"time"

	orderhooks2 "delivrio.io/go/schema/hooks/orderhooks"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Colli struct {
	ent.Schema
}

func (Colli) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Colli) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", Address.Type).
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("sender", Address.Type).
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("parcel_shop", ParcelShop.Type).
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("click_collect_location", Location.Type).
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("order_lines", OrderLine.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("delivery_option", DeliveryOption.Type).
			Unique(),
		edge.To("document_file", DocumentFile.Type).
			Comment("Stores packing slips for quick printing. Carrier labels are attached to the shipment."),
		edge.To("shipment_parcel", ShipmentParcel.Type).
			Unique().
			Comment("A colli may only have 1 active shipment, cancelled shipments are moved to the other edge"),
		edge.To("cancelled_shipment_parcel", ShipmentParcel.Type).
			Comment("A ref to all cancelled shipments"),
		edge.From("order", Order.Type).
			Unique().
			Required().
			Ref("colli"),
		edge.To("packaging", Packaging.Type).
			Unique().
			Comment("Allows packaging to be predefined for this colli and will be used for the shipment parcel"),
		edge.From("print_job", PrintJob.Type).
			Ref("colli"),
	}
}

func (Colli) Hooks() []ent.Hook {
	return []ent.Hook{
		orderhooks2.CreateColliBarcode(),
		history.ColliCreate(),
		history.ColliUpdate(),
		orderhooks2.DeleteColli(),
		orderhooks2.UpdateOrderStatusOnColliMutate(),
		orderhooks2.UpdateColliClearDocs(),
	}
}

func (Colli) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("colli")),
	}
}

func (Colli) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("internal_barcode", "tenant_id").
			Unique(),
	}
}

func (Colli) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_barcode").
			Comment("Code128 type C compatible for faster reads").
			Positive().
			Optional(). // To Support DB auto-increment when not on SQLite
			Min(1_000_000),
		field.Enum("status").
			Values("Pending", "Dispatched", "Cancelled").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Enum("slip_print_status").
			Default("pending").
			Values("pending", "printed"),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("email_packing_slip_printed_at").
			Optional().
			Comment("When filled, the packing slip email has been fired").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Time("email_label_printed_at").
			Optional().
			Comment("When filled, the packing slip email has been fired. Consider moving to shipping parcel? There are trade offs").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}
