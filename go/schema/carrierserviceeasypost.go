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

type CarrierServiceEasyPost struct {
	ent.Schema
}

func (CarrierServiceEasyPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

func (CarrierServiceEasyPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier_service", CarrierService.Type).
			Ref("carrier_serv_easy_post").
			Unique().
			Required(),
		edge.To("carrier_add_serv_easy_post", CarrierAdditionalServiceEasyPost.Type),
	}
}

func (CarrierServiceEasyPost) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierServiceEasyPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_service_easy_post")),
	}
}

func (CarrierServiceEasyPost) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("api_key").Values(
			"First",
			"Priority",
			"Express",
			"GroundAdvantage",
			"LibraryMail",
			"MediaMail",
			"FirstClassMailInternational",
			"FirstClassPackageInternationalService",
			"PriorityMailInternational",
			"ExpressMailInternational",
		).Annotations(entgql.Skip()),
	}
}
