package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderHistory holds the schema definition for the OrderHistory entity.
type OrderHistory struct {
	ent.Schema
}

// Fields of the OrderHistory.
func (OrderHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			Immutable(),
		field.Enum("type").
			Immutable().
			Values("create", "update", "delete", "notify"),
	}
}

// Edges of the OrderHistory.
func (OrderHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("order_history").
			Unique().
			Required(),
	}
}

func (OrderHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ChangeHistoryEntityMixin("order"),
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("order_history")),
	}
}
