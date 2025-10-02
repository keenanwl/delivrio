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

type CarrierAdditionalServicePostNord struct {
	ent.Schema
}

func (CarrierAdditionalServicePostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierAdditionalServicePostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_post_nord", CarrierServicePostNord.Type).
			Ref("carrier_add_serv_post_nord").
			Unique(),
		edge.From("delivery_option_post_nord", DeliveryOptionPostNord.Type).
			Ref("carrier_add_serv_post_nord"),
		edge.To("countries_consignee", Country.Type),
		edge.To("countries_consignor", Country.Type),
	}
}

func (CarrierAdditionalServicePostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServicePostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_post_nord")),
	}
}

func (CarrierAdditionalServicePostNord) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("carrier_service_post_nord").Fields("internal_id").Unique(),
	}
}

func (CarrierAdditionalServicePostNord) Fields() []ent.Field {
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
		field.String("api_code").
			Comment("2 characters code identifying the additional service in the API request"),
	}
}
