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

type HypothesisTestDeliveryOption struct {
	ent.Schema
}

func (HypothesisTestDeliveryOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (HypothesisTestDeliveryOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hypothesis_test", HypothesisTest.Type).
			Unique().
			Required().
			Ref("hypothesis_test_delivery_option"),
		edge.From("hypothesis_test_delivery_option_request", HypothesisTestDeliveryOptionRequest.Type).
			Ref("hypothesis_test_delivery_option"),
		edge.To("delivery_option_group_one", DeliveryOption.Type),
		edge.To("delivery_option_group_two", DeliveryOption.Type),
	}
}

func (HypothesisTestDeliveryOption) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (HypothesisTestDeliveryOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("hypothesis_test_delivery_option")),
	}
}

func (HypothesisTestDeliveryOption) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("randomize_within_group_sort").
			Default(false),
		field.Bool("by_interval_rotation").
			Default(false),
		field.Int("rotation_interval_hours").
			Default(6),
		field.Bool("by_order").
			Default(false),
	}
}
