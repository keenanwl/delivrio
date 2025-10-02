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

type CarrierUSPS struct {
	ent.Schema
}

func (CarrierUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (CarrierUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("carrier", Carrier.Type).
			Ref("carrier_usps").
			Unique().
			Required(),
	}
}

func (CarrierUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (CarrierUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("carrier_usps")),
	}
}

func (CarrierUSPS) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_test_api").
			Default(false),
		field.String("consumer_key").
			Optional(),
		field.String("consumer_secret").
			Optional(),
		field.String("mid").
			Optional(),
		field.String("manifest_mid").
			Optional(),
		field.String("crid").
			Optional(),
		field.String("eps_account_number").
			Optional(),
	}
}
