package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type CarrierServiceDSV struct {
	ent.Schema
}

func (CarrierServiceDSV) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierServiceDSV) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_dsv").
			Unique().
			Required(),
		edge.To("carrier_additional_service_dsv", CarrierAdditionalServiceDSV.Type),
	}
}

func (CarrierServiceDSV) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceDSV) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_dsv")),
	}
}

func (CarrierServiceDSV) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
