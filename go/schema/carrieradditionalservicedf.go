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

type CarrierAdditionalServiceDF struct {
	ent.Schema
}

func (CarrierAdditionalServiceDF) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierAdditionalServiceDF) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_df", CarrierServiceDF.Type).
			Ref("carrier_additional_service_df"),
		edge.From("delivery_option_df", DeliveryOptionDF.Type).
			Ref("carrier_additional_service_df"),
	}
}

func (CarrierAdditionalServiceDF) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceDF) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_df")),
	}
}

func (CarrierAdditionalServiceDF) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("api_code"),
	}
}
