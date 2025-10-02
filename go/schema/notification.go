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

type Notification struct {
	ent.Schema
}

func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connection", Connection.Type).
			Required().
			Unique(),
		edge.To("email_template", EmailTemplate.Type).
			Required().
			Unique(),
	}
}

func (Notification) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("notification")),
	}
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("active").
			Default(true),
	}
}
