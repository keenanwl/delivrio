package schema

import (
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeliveryOptionPostNord struct {
	ent.Schema
}

func (DeliveryOptionPostNord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryOptionPostNord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_option_post_nord").
			Unique().
			Required(),
		edge.To("carrier_add_serv_post_nord", CarrierAdditionalServicePostNord.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)).
			Comment("The lookup is done via the internalID, so edge input not included here. Consider refactoring to a generic entity on top of the PN entity."),
	}
}

func (DeliveryOptionPostNord) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (DeliveryOptionPostNord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_option_post_nord")),
	}
}

func (DeliveryOptionPostNord) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("format_zpl").
			Default(true),
	}
}
