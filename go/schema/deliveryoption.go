package schema

import (
	"delivrio.io/go/schema/hooks/deliveryoptionhooks"
	"delivrio.io/go/schema/mixins"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeliveryOption struct {
	ent.Schema
}

func (DeliveryOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier", Carrier.Type).
			Required().
			Unique(),
		edge.To("delivery_rule", DeliveryRule.Type),
		edge.To("delivery_option_dao", DeliveryOptionDAO.Type).
			Unique(),
		edge.To("delivery_option_df", DeliveryOptionDF.Type).
			Unique(),
		edge.To("delivery_option_dsv", DeliveryOptionDSV.Type).
			Unique(),
		edge.To("delivery_option_easy_post", DeliveryOptionEasyPost.Type).
			Unique(),
		edge.To("delivery_option_gls", DeliveryOptionGLS.Type).
			Unique(),
		edge.To("delivery_option_post_nord", DeliveryOptionPostNord.Type).
			Unique(),
		edge.To("delivery_option_usps", DeliveryOptionUSPS.Type).
			Unique(),
		edge.To("delivery_option_bring", DeliveryOptionBring.Type).
			Unique(),

		edge.From("return_portals", ReturnPortal.Type).
			Ref("delivery_options"),
		edge.From("colli", Colli.Type).
			Ref("delivery_option"),
		edge.From("return_colli", ReturnColli.Type).
			Ref("delivery_option"),
		edge.To("carrier_service", CarrierService.Type).
			Unique().
			Required(),
		edge.From("connection", Connection.Type).
			Ref("delivery_option").
			Unique().
			Required(),
		edge.From("connection_default", Connection.Type).
			Ref("default_delivery_option").
			Comment("The default delivery option for the connection. Unique since DO is already pinned to a single connection.").
			Unique(),
		edge.From("hypothesis_test_delivery_option_group_one", HypothesisTestDeliveryOption.Type).
			Ref("delivery_option_group_one"),
		edge.From("hypothesis_test_delivery_option_group_two", HypothesisTestDeliveryOption.Type).
			Ref("delivery_option_group_two"),
		edge.From("hypothesis_test_delivery_option_lookup", HypothesisTestDeliveryOptionLookup.Type).
			Ref("delivery_option"),
		edge.To("click_collect_location", Location.Type),
		edge.To("email_click_collect_at_store", EmailTemplate.Type).
			Unique(),

		edge.From("consolidation", Consolidation.Type).
			Ref("delivery_option"),
		edge.To("default_packaging", Packaging.Type).
			Unique(),
	}
}

func (DeliveryOption) Hooks() []ent.Hook {
	return []ent.Hook{
		deliveryoptionhooks.PreventConflictingIntegrations(),
	}
}

func (DeliveryOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option")),
		mixins.ArchiveMixin{},
	}
}
func (DeliveryOption) Indexes() []ent.Index {
	return []ent.Index{}
}

func (DeliveryOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("sort_order"),
		field.Int("click_option_display_count").
			Optional().
			Default(3).
			Max(20).
			Min(1),
		field.String("description").
			Optional(),
		field.Bool("click_collect").
			Optional().
			Default(false),
		field.Bool("override_sender_address").
			Optional().
			Default(false),
		field.Bool("override_return_address").
			Optional().
			Default(false),
		field.Bool("hide_delivery_option").
			Optional().
			Default(false),
		field.Int("delivery_estimate_from").
			Optional(),
		field.Int("delivery_estimate_to").
			Optional(),
		field.Bool("webshipper_integration").
			Default(false),
		field.Int("webshipper_id").
			Optional().
			Default(1).
			Positive(),
		field.Bool("shipmondo_integration").
			Default(false),
		field.String("shipmondo_delivery_option").
			Comment("May contain placeholders").
			Optional(),
		field.Bool("customs_enabled").
			Comment("Since some services are customs optional").
			Default(false),
		field.String("customs_signer").
			Comment("Who is responsible for signing of custom docs").
			Optional(),
		field.Bool("hide_if_company_empty").
			Default(false).
			Comment("Toggle to hide this rate if a company field is not provided."),
	}
}
