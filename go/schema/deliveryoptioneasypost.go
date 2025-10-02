package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type DeliveryOptionEasyPost struct {
	ent.Schema
}

func (DeliveryOptionEasyPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryOptionEasyPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_easy_post").
			Unique().
			Required(),
		edge.To("carrier_add_serv_easy_post", CarrierAdditionalServiceEasyPost.Type),
	}
}

func (DeliveryOptionEasyPost) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DeliveryOptionEasyPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_easy_post")),
	}
}

func (DeliveryOptionEasyPost) Fields() []ent.Field {
	return []ent.Field{
		//
	}
}
