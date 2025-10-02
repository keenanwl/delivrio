package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type WorkspaceRecentScan struct {
	ent.Schema
}

func (WorkspaceRecentScan) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (WorkspaceRecentScan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shipment_parcel", ShipmentParcel.Type).
			Unique(),
		edge.To("user", User.Type).
			Unique().
			Required().
			Immutable(),
	}
}

func (WorkspaceRecentScan) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (WorkspaceRecentScan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("workspace_recent_scan")),
	}
}

func (WorkspaceRecentScan) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}
