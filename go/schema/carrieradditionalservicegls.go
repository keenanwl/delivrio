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

type CarrierAdditionalServiceGLS struct {
	ent.Schema
}

func (CarrierAdditionalServiceGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierAdditionalServiceGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_gls", CarrierServiceGLS.Type).
			Ref("carrier_additional_service_gls").
			Unique(),
		edge.From("delivery_option_gls", DeliveryOptionGLS.Type).
			Ref("carrier_additional_service_gls"),
		edge.To("countries_consignee", Country.Type),
		edge.To("countries_consignor", Country.Type),
	}
}

func (CarrierAdditionalServiceGLS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_gls")),
	}
}

func (CarrierAdditionalServiceGLS) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.Bool("mandatory").
			Comment(""),
		field.Bool("all_countries_consignor").
			Comment("When false, only edge countries will validate on this consignor service").
			Default(false),
		field.Bool("all_countries_consignee").
			Comment("When false, only edge countries will validate on this consignee service").
			Default(false),
		field.String("internal_id").
			Comment(""),
	}
}
