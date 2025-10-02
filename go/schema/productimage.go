package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ProductImage struct {
	ent.Schema
}

func (ProductImage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (ProductImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).
			Unique().
			Required(),
		edge.To("product_variant", ProductVariant.Type),
	}
}

func (ProductImage) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ProductImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("product_image")),
	}
}

func (ProductImage) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_id").
			Optional().
			Nillable(),
		field.String("url").
			NotEmpty(),
	}
}

func (ProductImage) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("tenant").
			Fields("external_id").
			Unique(),
	}
}
