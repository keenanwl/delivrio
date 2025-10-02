package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type ConnectionLookup struct {
	ent.Schema
}

func (ConnectionLookup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ConnectionLookup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connections", Connection.Type).
			Unique().
			Immutable(),
	}
}

func (ConnectionLookup) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ConnectionLookup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connection_lookup")),
	}
}

func (ConnectionLookup) Fields() []ent.Field {
	return []ent.Field{
		field.String("payload"),
		field.Int("options_output_count"),
		field.String("error").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			).
			Immutable(),
	}
}
