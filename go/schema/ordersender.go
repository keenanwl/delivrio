package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type OrderSender struct {
	ent.Schema
}

func (OrderSender) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (OrderSender) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OrderSender) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (OrderSender) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("order_sender")),
	}
}

func (OrderSender) Fields() []ent.Field {
	return []ent.Field{
		field.String("uniqueness_id").Unique().Optional().Annotations(entgql.Skip(entgql.SkipAll)),
		field.String("first_name"),
		field.String("last_name"),
		field.String("email"),
		field.String("phone_number"),
		field.String("vat_number"),
	}
}
