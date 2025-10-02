package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Consolidation struct {
	ent.Schema
}

func (Consolidation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Consolidation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pallets", Pallet.Type),
		edge.To("orders", Order.Type),
		edge.To("delivery_option", DeliveryOption.Type).
			Unique(),
		edge.To("recipient", Address.Type).
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("sender", Address.Type).
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.From("shipment", Shipment.Type).
			Unique().
			Ref("consolidation"),
		edge.From("cancelled_shipments", Shipment.Type).
			Ref("old_consolidation"),
	}
}

func (Consolidation) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Consolidation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("consolidation")),
	}
}

func (Consolidation) Fields() []ent.Field {
	return []ent.Field{
		field.String("public_id"),
		field.String("description").
			Optional(),
		field.Enum("status").
			Default("Pending").
			Values("Pending", "Prebooked", "Booked", "Cancelled"),
		field.Time("created_at").
			Optional().
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			).
			Immutable(),
	}
}
