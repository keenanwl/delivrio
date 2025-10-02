package schema

import (
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserSeat holds the schema definition for the UserSeat entity.
type UserSeat struct {
	ent.Schema
}

func (UserSeat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the UserSeat.
func (UserSeat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		field.String("surname").Optional(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Sensitive().Annotations(entgql.Skip()),
		field.String("hash").Sensitive().Annotations(entgql.Skip()),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the UserSeat.
func (UserSeat) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (UserSeat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("user_seat")),
	}
}
