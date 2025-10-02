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

type DeliveryOptionUSPS struct {
	ent.Schema
}

func (DeliveryOptionUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryOptionUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_usps").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("carrier_additional_service_usps", CarrierAdditionalServiceUSPS.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (DeliveryOptionUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DeliveryOptionUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_usps")),
	}
}

func (DeliveryOptionUSPS) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("format_zpl").
			Default(true),
	}
}
