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

type ConnectionBrand struct {
	ent.Schema
}

func (ConnectionBrand) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ConnectionBrand) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").
			Unique(),
		field.Enum("internal_id").
			Values("shopify").
			Default("shopify"),
		field.String("logo_url").
			Optional(),
	}
}

func (ConnectionBrand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connection", Connection.Type).
			Ref("connection_brand"),
	}
}

func (ConnectionBrand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connection_brand")),
	}
}
