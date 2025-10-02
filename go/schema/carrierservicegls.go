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

type CarrierServiceGLS struct {
	ent.Schema
}

func (CarrierServiceGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

func (CarrierServiceGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_gls").
			Unique().
			Required(),
		edge.To("carrier_additional_service_gls", CarrierAdditionalServiceGLS.Type),
	}
}

func (CarrierServiceGLS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_gls")),
	}
}

func (CarrierServiceGLS) Fields() []ent.Field {
	return []ent.Field{
		field.String("api_key").
			Unique().
			Nillable().
			Optional(),
		field.Enum("api_value").
			Values("Y", "numeric_string", "none"),
	}
}
