package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessRight holds the schema definition for the AccessRight entity.
type AccessRight struct {
	ent.Schema
}

func (AccessRight) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// Fields of the AccessRight.
func (AccessRight) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Unique(),
		field.String("internal_id").Unique(),
	}
}

// Edges of the AccessRight.
func (AccessRight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assigned_access_right", SeatGroup.Type).
			Ref("assigned_access_right").
			Through("seat_group_access_right", SeatGroupAccessRight.Type),
	}
}

func (AccessRight) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("access_right")),
	}
}
