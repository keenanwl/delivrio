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

type CarrierService struct {
	ent.Schema
}

func (CarrierService) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierService) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("internal_id").
			Unique(),
		field.Bool("return").
			Default(false),
		field.Bool("consolidation").
			Default(false),
		field.Bool("delivery_point_optional").
			Default(false),
		field.Bool("delivery_point_required").
			Default(false),
	}
}

func (CarrierService) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier_service_post_nord", CarrierServicePostNord.Type).
			Unique(),
		edge.To("carrier_service_dao", CarrierServiceDAO.Type).
			Unique(),
		edge.To("carrier_service_df", CarrierServiceDF.Type).
			Unique(),
		edge.To("carrier_service_dsv", CarrierServiceDSV.Type).
			Unique(),
		edge.To("carrier_serv_easy_post", CarrierServiceEasyPost.Type).
			Unique(),
		edge.To("carrier_service_gls", CarrierServiceGLS.Type).
			Unique(),
		edge.To("carrier_service_usps", CarrierServiceUSPS.Type).
			Unique(),
		edge.To("carrier_service_bring", CarrierServiceBring.Type).
			Unique(),

		edge.From("carrier_brand", CarrierBrand.Type).
			Ref("carrier_service").
			Unique().
			Required(),
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("carrier_service"),
	}
}

func (CarrierService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service")),
	}
}
