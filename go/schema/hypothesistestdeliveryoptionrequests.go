package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type HypothesisTestDeliveryOptionRequest struct {
	ent.Schema
}

func (HypothesisTestDeliveryOptionRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (HypothesisTestDeliveryOptionRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hypothesis_test_delivery_option", HypothesisTestDeliveryOption.Type).
			Required().
			Unique(),
		edge.To("order", Order.Type).
			Unique(),
		edge.To("hypothesis_test_delivery_option_lookup", HypothesisTestDeliveryOptionLookup.Type),
	}
}

func (HypothesisTestDeliveryOptionRequest) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (HypothesisTestDeliveryOptionRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("hypothesis_test_delivery_option_request")),
	}
}

func (HypothesisTestDeliveryOptionRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_hash"),
		field.String("shipping_address_hash"),
		field.Bool("is_control_group"),
		field.Uint("request_count"),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("last_requested_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).UpdateDefault(time.Now),
	}
}
