package schema

import (
	"time"

	"delivrio.io/go/schema/hooks/systemeventhooks"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.String("title"),
		field.String("body_html").
			Optional(),
		field.Enum("status").
			Default("active").
			Values("active", "archived", "draft"),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (Product) Hooks() []ent.Hook {
	return []ent.Hook{
		systemeventhooks.CreateProduct(),
		systemeventhooks.UpdateProduct(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product_tags", ProductTag.Type).
			Ref("products"),
		edge.To("product_variant", ProductVariant.Type),
		edge.From("product_image", ProductImage.Type).
			Ref("product"),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("product")),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("tenant").Fields("external_id").Unique(),
	}
}
