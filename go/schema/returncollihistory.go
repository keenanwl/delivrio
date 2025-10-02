package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ReturnColliHistory struct {
	ent.Schema
}

func (ReturnColliHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ReturnColliHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("return_colli", ReturnColli.Type).
			Ref("return_colli_history").
			Unique().
			Required(),
	}
}

func (ReturnColliHistory) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ReturnColliHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ChangeHistoryEntityMixin("return_colli"),
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("return_colli_history")),
	}
}

func (ReturnColliHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			Immutable(),
		field.Enum("type").
			Immutable().
			Values("create", "update", "delete", "notify"),
	}
}
