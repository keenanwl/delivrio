package schema

import (
	"time"

	"delivrio.io/go/schema/hooks/emailtemplatevalidators"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EmailTemplate struct {
	ent.Schema
}

func (EmailTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (EmailTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("return_portal_confirmation_label", ReturnPortal.Type).
			Ref("email_confirmation_label"),
		edge.From("return_portal_confirmation_qr_code", ReturnPortal.Type).
			Ref("email_confirmation_qr_code"),
		edge.From("return_portal_received", ReturnPortal.Type).
			Ref("email_received"),
		edge.From("return_portal_accepted", ReturnPortal.Type).
			Ref("email_accepted"),
		edge.From("delivery_option_click_collect_at_store", DeliveryOption.Type).
			Ref("email_click_collect_at_store"),
		edge.From("notifications", Notification.Type).
			Ref("email_template"),
	}
}

func (EmailTemplate) Hooks() []ent.Hook {
	return []ent.Hook{
		emailtemplatevalidators.CreateUpdateEmailTemplate(),
	}
}

func (EmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("email_template")),
	}
}

func (EmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(64),
		field.String("subject").
			MaxLen(255).
			Default(""),
		field.String("html_template").
			Default(""),
		field.Enum("merge_type").
			Values(
				"return_colli_label",
				"return_colli_qr",
				"return_colli_received",
				"return_colli_accepted",
				"order_confirmation",
				"order_picked",
			).
			Default("return_colli_label"),
		field.Time("created_at").
			Default(time.Now).
			Optional().
			Immutable().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}
