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

type CarrierEasyPost struct {
	ent.Schema
}

func (CarrierEasyPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierEasyPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_easy_post").
			Unique().
			Required(),
	}
}

func (CarrierEasyPost) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierEasyPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_easy_post")),
	}
}

func (CarrierEasyPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("api_key"),
		field.Bool("test").
			Default(true),
		field.Strings("carrier_accounts").
			Comment("When > 1, then we use rate, then buy. =1 one-call buy. Former not implemented in first round.").
			Default([]string{}),
	}
}
