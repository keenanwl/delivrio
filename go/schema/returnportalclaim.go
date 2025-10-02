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

type ReturnPortalClaim struct {
	ent.Schema
}

func (ReturnPortalClaim) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ReturnPortalClaim) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("return_portal", ReturnPortal.Type).
			Ref("return_portal_claim").
			Unique().
			Required(),
		edge.To("return_location", Location.Type).
			Unique().
			Comment("Return to address"),
		edge.From("return_order_line", ReturnOrderLine.Type).
			Ref("return_portal_claim"),
	}
}

func (ReturnPortalClaim) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ReturnPortalClaim) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("return_portal_claim")),
	}
}

func (ReturnPortalClaim) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Bool("restockable"),
		field.Bool("archived").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}
