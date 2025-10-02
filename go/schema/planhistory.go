package schema

import (
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlanHistory holds the schema definition for the PlanHistory entity.
type PlanHistory struct {
	ent.Schema
}

// Fields of the PlanHistory.
func (PlanHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the PlanHistory.
func (PlanHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("changed_by", User.Type).
			Ref("plan_history_user").
			Unique().
			Required(),
		edge.From("changed_from", Plan.Type).
			Ref("plan_history_plan").
			Unique().
			Required(),
	}
}

// Mixin of the Tenant schema.
func (PlanHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ChangeHistoryEntityMixin("plan"),
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("plan_history")),
	}
}
