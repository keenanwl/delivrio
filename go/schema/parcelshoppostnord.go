package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ParcelShopPostNord struct {
	ent.Schema
}

func (ParcelShopPostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ParcelShopPostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parcel_shop", ParcelShop.Type).
			Required().
			Unique().
			Ref("parcel_shop_post_nord"),
		edge.To("address_delivery", AddressGlobal.Type).
			Required().
			Unique(),
	}
}

func (ParcelShopPostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}
func (ParcelShopPostNord) Indexes() []ent.Index {
	return []ent.Index{
		//
	}
}

func (ParcelShopPostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("parcel_shop_post_nord")),
	}
}

func (ParcelShopPostNord) Fields() []ent.Field {
	return []ent.Field{
		field.String("service_point_id"),
		field.String("pudoid").Unique(),
		field.String("type_id").
			Comment("No idea what the options are aside from the default: 156. Maybe box, shop, etc??"),
	}
}
