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

type InventoryItem struct {
	ent.Schema
}

func (InventoryItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (InventoryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("country_harmonized_code", CountryHarmonizedCode.Type).
			Comment("Takes precedent over general HS code"),
		edge.To("country_of_origin", Country.Type).
			Comment("Can be null in Shopify").
			Unique(),
		edge.From("product_variant", ProductVariant.Type).
			Ref("inventory_item").
			Unique().
			Required(),
	}
}

func (InventoryItem) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (InventoryItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("inventory_item")),
	}
}

func (InventoryItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("external_ID").
			Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.String("code").
			Optional().
			Nillable().
			Comment("Used when country specific code not available"),
		field.String("sku").
			Optional().
			Nillable().
			Comment("Duplicated to match Shopify InventoryItem/ProductVariant"),
	}
}
