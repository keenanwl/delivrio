package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type DeliveryOptionDAO struct {
	ent.Schema
}

func (DeliveryOptionDAO) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (DeliveryOptionDAO) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_dao").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("carrier_additional_service_dao", CarrierAdditionalServiceDAO.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (DeliveryOptionDAO) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DeliveryOptionDAO) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_dao")),
	}
}

func (DeliveryOptionDAO) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
