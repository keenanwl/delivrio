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

// Carrier holds the schema definition for the Carrier entity.
type Carrier struct {
	ent.Schema
}

func (Carrier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (Carrier) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("sync_cancelation").
			Default(false),
	}
}

func (Carrier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier_brand", CarrierBrand.Type).
			Required().
			Unique(),
		edge.To("carrier_dao", CarrierDAO.Type).
			Unique(),
		edge.To("carrier_df", CarrierDF.Type).
			Unique(),
		edge.To("carrier_dsv", CarrierDSV.Type).
			Unique(),
		edge.To("carrier_easy_post", CarrierEasyPost.Type).
			Unique(),
		edge.To("carrier_gls", CarrierGLS.Type).
			Unique(),
		edge.To("carrier_post_nord", CarrierPostNord.Type).
			Unique(),
		edge.To("carrier_usps", CarrierUSPS.Type).
			Unique(),
		edge.To("carrier_bring", CarrierBring.Type).
			Unique(),
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("carrier"),
		edge.From("shipment", Shipment.Type).
			Ref("carrier"),
	}
}

func (Carrier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier")),
	}
}
