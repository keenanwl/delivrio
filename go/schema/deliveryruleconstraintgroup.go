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

type DeliveryRuleConstraintGroup struct {
	ent.Schema
}

func (DeliveryRuleConstraintGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryRuleConstraintGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("constraint_logic").
			Default("and").
			Values("and", "or"),
	}
}

func (DeliveryRuleConstraintGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("delivery_rule_constraints", DeliveryRuleConstraint.Type),
		edge.From("delivery_rule", DeliveryRule.Type).
			Ref("delivery_rule_constraint_group").
			Unique().
			Required(),
	}
}

func (DeliveryRuleConstraintGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_rule_constraint_group")),
	}
}
