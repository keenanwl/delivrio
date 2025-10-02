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

type CarrierAdditionalServiceBring struct {
	ent.Schema
}

func (CarrierAdditionalServiceBring) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierAdditionalServiceBring) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_bring", CarrierServiceBring.Type).
			Ref("carrier_additional_service_bring").
			Unique(),
		edge.From("delivery_option_bring", DeliveryOptionBring.Type).
			Ref("carrier_additional_service_bring"),
	}
}

func (CarrierAdditionalServiceBring) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceBring) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_bring")),
	}
}

func (CarrierAdditionalServiceBring) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("api_code_booking"),
	}
}
