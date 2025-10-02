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

// SeatGroupAccessRight holds the schema definition for the SeatGroupAccessRight entity.
type SeatGroupAccessRight struct {
	ent.Schema
}

func (SeatGroupAccessRight) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the SeatGroupAccessRight.
func (SeatGroupAccessRight) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("level").Default("none").Values("none", "read", "write"),
		field.String("access_right_id").
			GoType(pulid.ID("")),
		field.String("seat_group_id").
			GoType(pulid.ID("")),
	}
}

// Edges of the SeatGroupAccessRight.
func (SeatGroupAccessRight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("access_right", AccessRight.Type).
			Unique().
			Required().
			Field("access_right_id"),
		edge.To("seat_group", SeatGroup.Type).
			Unique().
			Required().
			Field("seat_group_id"),
	}
}

func (SeatGroupAccessRight) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("seat_group_access_right")),
	}
}
