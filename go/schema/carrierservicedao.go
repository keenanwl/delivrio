package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type CarrierServiceDAO struct {
	ent.Schema
}

func (CarrierServiceDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierServiceDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_service_dao").
			Unique().
			Required(),
		edge.To("carrier_additional_service_dao", CarrierAdditionalServiceDAO.Type),
	}
}

func (CarrierServiceDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_dao")),
	}
}

func (CarrierServiceDAO) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
