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
)

type ParcelShop struct {
	ent.Schema
}

func (ParcelShop) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ParcelShop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parcel_shop_dao", ParcelShopDAO.Type).
			Unique(),
		edge.To("parcel_shop_post_nord", ParcelShopPostNord.Type).
			Unique(),
		edge.To("parcel_shop_gls", ParcelShopGLS.Type).
			Unique(),
		edge.To("parcel_shop_bring", ParcelShopBring.Type).
			Unique(),
		edge.To("carrier_brand", CarrierBrand.Type).
			Required().
			Unique(),
		edge.To("address", AddressGlobal.Type).
			Required().
			Unique(),
		edge.From("colli", Colli.Type).
			Ref("parcel_shop"),
		edge.To("business_hours_period", BusinessHoursPeriod.Type),
	}
}

func (ParcelShop) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ParcelShop) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("parcel_shop")),
	}
}

func (ParcelShop) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("last_updated").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}
