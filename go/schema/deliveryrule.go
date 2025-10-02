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

type DeliveryRule struct {
	ent.Schema
}

func (DeliveryRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("price").
			Default(20.00),
	}
}

func (DeliveryRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("delivery_rule_constraint_group", DeliveryRuleConstraintGroup.Type).
			Comment("Since constraint groups can be && or ||, we need to have groups of multiple constraints"),
		// TODO refactor to use required
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("delivery_rule").
			Unique(),
		edge.From("country", Country.Type).
			Ref("delivery_rule"),
		edge.To("currency", Currency.Type).
			//Required(). // Enable after next deploy
			Unique(),
	}
}

func (DeliveryRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_rule")),
	}
}
