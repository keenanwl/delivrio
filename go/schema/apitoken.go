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

type APIToken struct {
	ent.Schema
}

func (APIToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (APIToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("api_token").
			Unique().
			Required(),
	}
}

func (APIToken) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (APIToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("api_token")),
	}
}

func (APIToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("User supplied name for this token"),
		field.String("hashed_token").
			Unique().
			Immutable().
			MinLen(60).
			NotEmpty().
			Sensitive(),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("last_used").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}
