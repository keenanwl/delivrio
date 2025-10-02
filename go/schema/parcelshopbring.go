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

type ParcelShopBring struct {
	ent.Schema
}

func (ParcelShopBring) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ParcelShopBring) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop", ParcelShop.Type).
			Required().
			Unique().
			Ref("parcel_shop_bring"),
		edge.To("address_delivery", AddressGlobal.Type).
			Required().
			Unique(),
	}
}

func (ParcelShopBring) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ParcelShopBring) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("parcel_shop_bring")),
	}
}

func (ParcelShopBring) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("point_type").
			// Non-numeric since Graphql complains otherwise
			Values(
				"one",
				"four",
				"nineteen",
				"twenty_one",
				"thirty_two",
				"thirty_four",
				"thirty_seven",
				"thirty_eight",
				"thirty_nine",
				"eighty_five",
				"eighty_six",
				"SmartPOST",
				"Posti",
				"Noutopiste",
				"LOCKER",
				"Unknown",
			).
			Comment("https://developer.bring.com/api/pickup-point/#pickup-point-types"),
		field.String("bring_id").
			Unique(),
	}
}
