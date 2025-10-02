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

// ConnectOptionCarrier holds the schema definition for the ConnectOptionCarrier entity.
type ConnectOptionCarrier struct {
	ent.Schema
}

func (ConnectOptionCarrier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the ConnectOptionCarrier.
func (ConnectOptionCarrier) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the ConnectOptionCarrier.
func (ConnectOptionCarrier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("connect_option_carriers"),
	}
}

func (ConnectOptionCarrier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connect_option_carrier")),
	}
}
