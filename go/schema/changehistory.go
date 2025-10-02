package schema

import (
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChangeHistory holds the schema definition for the ChangeHistory entity.
type ChangeHistory struct {
	ent.Schema
}

func (ChangeHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// Fields of the ChangeHistory.
func (ChangeHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Immutable(),
		field.Enum("origin").
			Immutable().
			Default("unknown").
			Values("unknown", "background", "rest_api", "web_client", "print_client", "seed"),
	}
}

// Edges of the ChangeHistory.
func (ChangeHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plan_history", PlanHistory.Type),
		edge.To("user", User.Type).Unique(),
		edge.To("order_history", OrderHistory.Type),
		edge.To("shipment_history", ShipmentHistory.Type),
		edge.To("return_colli_history", ReturnColliHistory.Type),
	}
}

func (ChangeHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("change_history")),
	}
}
