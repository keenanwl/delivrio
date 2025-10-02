package schema

import (
	"entgo.io/ent/schema/index"
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SystemEvents holds the schema definition for the SystemEvents entity.
type SystemEvents struct {
	ent.Schema
}

func (SystemEvents) Annotations() []schema.Annotation {

	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// Fields of the SystemEvents.
func (SystemEvents) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("event_type").
			Values(
				"background_tasks",
				"shopify_product_sync",
				"shopify_order_sync",
				// Does not appear we can get a single list with two categories
				// So we run a separate sync-cancelled job
				"shopify_order_cancelled_sync",
				"background_product_mutate",
				"send_notifications",
				"sync_cancelled_shipments",
			),
		field.String("event_type_id").
			Optional(),
		field.Enum("status").
			Values("running", "fail", "success"),
		field.String("description"),
		field.String("data").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable(),
	}
}

// Edges of the SystemEvents.
func (SystemEvents) Edges() []ent.Edge {
	return nil
}

func (SystemEvents) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("system_events")),
	}
}

func (SystemEvents) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at"),
	}
}
