package schema

import (
	"delivrio.io/go/schema/hooks/hookgls"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type DeliveryOptionGLS struct {
	ent.Schema
}

func (DeliveryOptionGLS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (DeliveryOptionGLS) Fields() []ent.Field {
	return []ent.Field{}
}

func (DeliveryOptionGLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_gls").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("carrier_additional_service_gls", CarrierAdditionalServiceGLS.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)).
			Comment("The lookup is done via the internalID, so edge input not included here. Consider refactoring to a generic entity on top of the GLS entity."),
	}
}

func (DeliveryOptionGLS) Hooks() []ent.Hook {
	return []ent.Hook{
		hookgls.UpdateDeliveryOptionGLS(),
	}
}

func (DeliveryOptionGLS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_gls")),
	}
}
