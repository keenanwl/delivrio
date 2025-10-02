package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type CarrierServiceDF struct {
	ent.Schema
}

func (CarrierServiceDF) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierServiceDF) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_df").
			Unique().
			Required(),
		edge.To("carrier_additional_service_df", CarrierAdditionalServiceDF.Type),
	}
}

func (CarrierServiceDF) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceDF) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_df")),
	}
}

func (CarrierServiceDF) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
