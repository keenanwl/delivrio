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

type CarrierServiceUSPS struct {
	ent.Schema
}

func (CarrierServiceUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (CarrierServiceUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_usps").
			Unique().
			Required(),
		edge.To("carrier_additional_service_usps", CarrierAdditionalServiceUSPS.Type),
	}
}

func (CarrierServiceUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_usps")),
	}
}

func (CarrierServiceUSPS) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("api_key").Values(
			"PARCEL_SELECT",
			"PARCEL_SELECT_LIGHTWEIGHT",
			"USPS_CONNECT_LOCAL",
			"USPS_CONNECT_REGIONAL",
			"USPS_CONNECT_MAIL",
			"USPS_GROUND_ADVANTAGE",
			"PRIORITY_MAIL_EXPRESS",
			"PRIORITY_MAIL",
			"FIRST-CLASS_PACKAGE_SERVICE",
			"LIBRARY_MAIL",
			"MEDIA_MAIL",
			"BOUND_PRINTED_MATTER",
			"DOMESTIC_MATTER_FOR_THE_BLIND",
			// Returns
			"FIRST-CLASS_PACKAGE_RETURN_SERVICE",
			"GROUND_RETURN_SERVICE",
			"PRIORITY_MAIL_EXPRESS_RETURN_SERVICE",
			"PRIORITY_MAIL_RETURN_SERVICE",
			"USPS_GROUND_ADVANTAGE_RETURN_SERVICE",
		).Annotations(entgql.Skip()),
	}
}
