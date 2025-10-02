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

// ConnectOptionPlatform holds the schema definition for the ConnectOptionPlatform entity.
type ConnectOptionPlatform struct {
	ent.Schema
}

func (ConnectOptionPlatform) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the ConnectOptionPlatform.
func (ConnectOptionPlatform) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the ConnectOptionPlatform.
func (ConnectOptionPlatform) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("connect_option_platforms"),
	}
}

func (ConnectOptionPlatform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connect_option_platform")),
	}
}
