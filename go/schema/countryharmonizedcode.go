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

type CountryHarmonizedCode struct {
	ent.Schema
}

func (CountryHarmonizedCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CountryHarmonizedCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inventory_item", InventoryItem.Type).
			Ref("country_harmonized_code").
			Unique().
			Required(),
		edge.To("country", Country.Type).
			Unique().
			Required(),
	}
}

func (CountryHarmonizedCode) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CountryHarmonizedCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("country_harmonized_code")),
	}
}

func (CountryHarmonizedCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
	}
}
