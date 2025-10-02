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

type CarrierAdditionalServiceDSV struct {
	ent.Schema
}

func (CarrierAdditionalServiceDSV) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierAdditionalServiceDSV) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_dsv", CarrierServiceDSV.Type).
			Ref("carrier_additional_service_dsv"),
		edge.From("delivery_option_dsv", DeliveryOptionDSV.Type).
			Ref("carrier_additional_service_dsv"),
	}
}

func (CarrierAdditionalServiceDSV) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceDSV) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_dsv")),
	}
}

func (CarrierAdditionalServiceDSV) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("api_code"),
	}
}
