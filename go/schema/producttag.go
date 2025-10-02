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
	"entgo.io/ent/schema/index"
)

// ProductTag holds the schema definition for the ProductTag entity.
type ProductTag struct {
	ent.Schema
}

func (ProductTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the ProductTag.
func (ProductTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable(),
	}
}

// Edges of the ProductTag.
func (ProductTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}

func (ProductTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("product_tag")),
	}
}

func (ProductTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("tenant").Fields("name").Unique(),
	}
}
