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

type ParcelShopDAO struct {
	ent.Schema
}

func (ParcelShopDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ParcelShopDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop", ParcelShop.Type).
			Required().
			Unique().
			Ref("parcel_shop_dao"),
	}
}

func (ParcelShopDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ParcelShopDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("parcel_shop_dao")),
	}
}

func (ParcelShopDAO) Fields() []ent.Field {
	return []ent.Field{
		field.String("shop_id"),
	}
}
