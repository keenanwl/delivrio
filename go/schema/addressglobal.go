package schema

import (
	"delivrio.io/go/schema/delivrioannotations"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AddressGlobal no tenant: that is what makes this different than Address
type AddressGlobal struct {
	ent.Schema
}

func (AddressGlobal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		delivrioannotations.Clone(),
	}
}

func (AddressGlobal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop_post_nord_delivery", ParcelShopPostNord.Type).
			Ref("address_delivery").
			Unique(),
		edge.From("parcel_shop_bring_delivery", ParcelShopBring.Type).
			Ref("address_delivery").
			Unique(),
		edge.From("parcel_shop", ParcelShop.Type).
			Ref("address").
			Unique(),
		edge.To("country", Country.Type).
			Required().
			Unique(),
	}
}

func (AddressGlobal) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (AddressGlobal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("address_global")),
	}
}

func (AddressGlobal) Fields() []ent.Field {
	return []ent.Field{
		field.String("uniqueness_id").
			Unique().
			Optional().
			Annotations(entgql.Skip(entgql.SkipAll), delivrioannotations.Skip()),
		field.String("company").
			Optional(),
		field.String("address_one"),
		field.String("address_two").
			Optional(),
		field.String("city"),
		field.String("state").
			Optional(),
		field.String("zip"),
		field.Float("latitude").
			Default(0),
		field.Float("longitude").
			Default(0),
	}
}
