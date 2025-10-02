package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type CarrierBrand struct {
	ent.Schema
}

func (CarrierBrand) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierBrand) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("label_short").
			Comment("Accommodation for PostNord to become PN"),
		field.Enum("internal_id").Values(
			"bring",
			"dao",
			"df",
			"dsv",
			"easy_post",
			"gls",
			"dhl",
			"post_nord",
			"usps",
		),
		field.String("logo_url").Optional(),
		field.String("text_color").Optional().Default("#FFFFFF"),
		field.String("background_color").Optional().Default("#000000"),
	}
}

func (CarrierBrand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier_service", CarrierService.Type),
		edge.From("carrier", Carrier.Type).
			Ref("carrier_brand"),
		edge.From("parcel_shop", ParcelShop.Type).
			Ref("carrier_brand"),
		edge.From("packaging", Packaging.Type).
			Ref("carrier_brand"),
		edge.From("document", Document.Type).
			Ref("carrier_brand"),
	}
}

func (CarrierBrand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_brand")),
	}
}

func (CarrierBrand) Indexes() []ent.Index {
	return []ent.Index{
		// Enforce uniquness here so we have Enum types on
		// the internal_id field available client side
		index.Fields("internal_id").
			Unique(),
	}
}
