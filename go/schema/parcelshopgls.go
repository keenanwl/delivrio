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

type ParcelShopGLS struct {
	ent.Schema
}

func (ParcelShopGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ParcelShopGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop", ParcelShop.Type).
			Required().
			Unique().
			Ref("parcel_shop_gls"),
	}
}

func (ParcelShopGLS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ParcelShopGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("parcel_shop_gls")),
	}
}

func (ParcelShopGLS) Fields() []ent.Field {
	return []ent.Field{
		field.String("gls_parcel_shop_id").Unique(),
		field.String("partner_id").
			Optional().
			Comment("Only available in Group API"),
		field.String("type").
			Optional().
			Comment("Only available in Group API"),
	}
}
