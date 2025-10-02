package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

type HypothesisTestDeliveryOptionLookup struct {
	ent.Schema
}

func (HypothesisTestDeliveryOptionLookup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (HypothesisTestDeliveryOptionLookup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("delivery_option", DeliveryOption.Type).
			Unique().
			Required(),
		edge.From("hypothesis_test_delivery_option_request", HypothesisTestDeliveryOptionRequest.Type).
			Unique().
			Ref("hypothesis_test_delivery_option_lookup").
			Required(),
	}
}

func (HypothesisTestDeliveryOptionLookup) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("delivery_option", "hypothesis_test_delivery_option_request").
			Unique(),
	}
}

func (HypothesisTestDeliveryOptionLookup) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (HypothesisTestDeliveryOptionLookup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("hypothesis_test_delivery_option_lookup")),
	}
}

func (HypothesisTestDeliveryOptionLookup) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
