package schema

import (
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/schema/hooks/returncollihooks"
	"time"

	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ReturnColli struct {
	ent.Schema
}

func (ReturnColli) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (ReturnColli) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipient", Address.Type).
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("sender", Address.Type).
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.From("order", Order.Type).
			Required().
			Ref("return_colli").
			Unique(),
		edge.To("delivery_option", DeliveryOption.Type).
			Unique(),
		edge.To("return_portal", ReturnPortal.Type).
			Unique().
			Required(),
		edge.To("packaging", Packaging.Type).
			Unique().
			Comment("Allows packaging to be predefined for this colli and will be used for the return shipment parcel"),
		edge.To("return_order_line", ReturnOrderLine.Type),
		edge.To("return_colli_history", ReturnColliHistory.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (ReturnColli) Hooks() []ent.Hook {
	return []ent.Hook{
		returncollihooks.UpdateReturnColli(),
		history.ReturnColliCreate(),
		history.ReturnColliUpdate(),
	}
}

func (ReturnColli) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("return_colli")),
	}
}

func (ReturnColli) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expected_at").
			Nillable().
			Optional().
			Default(func() time.Time {
				// Just for demoing purposes
				now := time.Now()
				tomorrowMidday := time.Date(now.Year(), now.Month(), now.Day()+1, 12, 0, 0, 0, now.Location())
				return tomorrowMidday
			}),
		field.String("label_pdf").
			Optional(),
		field.String("label_png").
			Optional(),
		field.String("qr_code_png").
			Optional(),
		field.String("comment").
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Enum("status").
			Default("Opened").
			Values("Opened", "Pending", "Inbound", "Received", "Accepted", "Declined", "Deleted"),
		field.Time("email_received").
			Comment("Timestamp of email successfully sent after status changed to received").
			Nillable().
			Optional(),
		field.Time("email_accepted").
			Comment("Timestamp of email successfully sent after status changed to accepted").
			Nillable().
			Optional(),
		field.Time("email_confirmation_label").
			Comment("Timestamp of email successfully sent after status changed to pending").
			Nillable().
			Optional(),
		field.Time("email_confirmation_qr_code").
			Comment("Timestamp of email successfully sent after status changed to pending").
			Nillable().
			Optional(),
	}
}
