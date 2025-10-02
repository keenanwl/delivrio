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

type CarrierAdditionalServiceEasyPost struct {
	ent.Schema
}

func (CarrierAdditionalServiceEasyPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (CarrierAdditionalServiceEasyPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service_easy_post", CarrierServiceEasyPost.Type).
			Ref("carrier_add_serv_easy_post"),
		edge.From("delivery_option_easy_post", DeliveryOptionEasyPost.Type).
			Ref("carrier_add_serv_easy_post"),
	}
}

func (CarrierAdditionalServiceEasyPost) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierAdditionalServiceEasyPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_additional_service_easy_post")),
	}
}

func (CarrierAdditionalServiceEasyPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("label"),
		field.String("api_key"),
		field.String("api_value"),
	}
}
