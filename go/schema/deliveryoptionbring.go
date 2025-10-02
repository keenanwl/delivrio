package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeliveryOptionBring struct {
	ent.Schema
}

func (DeliveryOptionBring) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryOptionBring) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_bring").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("carrier_additional_service_bring", CarrierAdditionalServiceBring.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (DeliveryOptionBring) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DeliveryOptionBring) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_bring")),
	}
}

func (DeliveryOptionBring) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("electronic_customs").
			Default(false),
	}
}
