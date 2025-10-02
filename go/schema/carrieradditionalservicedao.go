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

type CarrierAdditionalServiceDAO struct {
	ent.Schema
}

func (CarrierAdditionalServiceDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierAdditionalServiceDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_dao", CarrierServiceDAO.Type).
			Ref("carrier_additional_service_dao"),
		edge.From("delivery_option_dao", DeliveryOptionDAO.Type).
			Ref("carrier_additional_service_dao"),
	}
}

func (CarrierAdditionalServiceDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_dao")),
	}
}

func (CarrierAdditionalServiceDAO) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("api_code"),
	}
}
