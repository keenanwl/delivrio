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

type ReturnPortal struct {
	ent.Schema
}

func (ReturnPortal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ReturnPortal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("return_portal_claim", ReturnPortalClaim.Type),
		edge.To("return_location", Location.Type),
		edge.To("delivery_options", DeliveryOption.Type),
		edge.To("connection", Connection.Type).
			Unique(),
		edge.To("email_confirmation_label", EmailTemplate.Type).
			Unique(),
		edge.To("email_confirmation_qr_code", EmailTemplate.Type).
			Unique(),
		edge.To("email_received", EmailTemplate.Type).
			Unique(),
		edge.To("email_accepted", EmailTemplate.Type).
			Unique(),
		edge.From("return_colli", ReturnColli.Type).
			Ref("return_portal"),
	}
}

func (ReturnPortal) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (ReturnPortal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("return_portal")),
	}
}

func (ReturnPortal) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("return_open_hours").
			Default(24 * 30),
		field.Bool("automatically_accept").
			Default(false),
	}
}
