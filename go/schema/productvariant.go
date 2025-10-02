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

// ProductVariant holds the schema definition for the ProductVariant entity.
type ProductVariant struct {
	ent.Schema
}

func (ProductVariant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the ProductVariant.
func (ProductVariant) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("archived").
			Default(false),
		field.String("external_ID").
			Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.String("description").
			Optional(),
		field.String("ean_number").
			Optional().
			Nillable(),
		field.Int("weight_g").
			Default(0).
			Optional().
			Nillable().
			NonNegative(),
		field.Int("dimension_length").
			Optional().
			Nillable().
			NonNegative(),
		field.Int("dimension_width").
			Optional().
			Nillable().
			NonNegative(),
		field.Int("dimension_height").
			Optional().
			Nillable().
			NonNegative(),
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

// TODO: add hook to prevent archiving the last variant

func (ProductVariant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("product_variant").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.From("order_lines", OrderLine.Type).
			Ref("product_variant"),
		edge.From("product_image", ProductImage.Type).
			Ref("product_variant"),
		edge.To("inventory_item", InventoryItem.Type).
			Unique(),
	}
}

func (ProductVariant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("product_variant")),
	}
}

func (ProductVariant) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("tenant").Fields("external_ID").Unique(),
	}
}
