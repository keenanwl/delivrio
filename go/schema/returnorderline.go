package schema

import (
	"delivrio.io/go/schema/hooks/returncollihooks"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ReturnOrderLine struct {
	ent.Schema
}

func (ReturnOrderLine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ReturnOrderLine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("return_colli", ReturnColli.Type).
			Ref("return_order_line").
			Required().
			Unique(),
		edge.To("order_line", OrderLine.Type).
			Required().
			Unique(),
		edge.To("return_portal_claim", ReturnPortalClaim.Type).
			Required().
			Unique(),
	}
}

func (ReturnOrderLine) Hooks() []ent.Hook {
	return []ent.Hook{
		returncollihooks.CreateReturnOrderLine(),
	}
}

func (ReturnOrderLine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("return_order_line")),
	}
}

func (ReturnOrderLine) Fields() []ent.Field {
	return []ent.Field{
		field.Int("units").Positive(),
	}
}
