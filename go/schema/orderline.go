package schema

import (
	"time"

	"delivrio.io/go/schema/hooks/orderhooks"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderLine struct {
	ent.Schema
}

func (OrderLine) Fields() []ent.Field {
	return []ent.Field{
		field.Float("unit_price"),
		field.Float("discount_allocation_amount").
			Comment("Amount removed from unit_price*units for customs docs"),
		field.String("external_id").
			Optional(),
		field.Int("units").Positive(),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("product_variant_id").
			GoType(pulid.ID("")),
		field.String("colli_id").
			GoType(pulid.ID("")),
	}
}

func (OrderLine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_variant", ProductVariant.Type).
			Unique().
			Required().
			Field("product_variant_id"),
		edge.From("colli", Colli.Type).
			Unique().
			Required().
			Ref("order_lines").
			Field("colli_id"),
		edge.From("return_order_line", ReturnOrderLine.Type).
			Ref("order_line"),
		edge.To("currency", Currency.Type).
			Unique().
			Required(),
	}
}

func (OrderLine) Hooks() []ent.Hook {
	return []ent.Hook{
		orderhooks.DeleteOrderLine(),
		orderhooks.UpdateOrderLineClearDocs(),
		orderhooks.CreateOrderLineClearDocs(),
	}
}

func (OrderLine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("order_line")),
	}
}
