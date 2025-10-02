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

type LocationTag struct {
	ent.Schema
}

func (LocationTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (LocationTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("location_tags"),
	}
}

func (LocationTag) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (LocationTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("location_tag")),
	}
}

func (LocationTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").Unique(),
		field.String("internal_id").Unique(),
	}
}
