package schema

import (
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/schema/hooks/constraint"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeliveryRuleConstraint struct {
	ent.Schema
}

func (DeliveryRuleConstraint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (DeliveryRuleConstraint) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("property_type").
			Values("total_weight", "cart_total", "day_of_week", "time_of_day", "product_tag", "all_products_tagged", "sku", "order_lines", "postal_code_numeric", "postal_code_string"),
		field.Enum("comparison").
			Values("equals", "not_equals", "between", "outside", "less_than", "greater_than", "contains", "prefix", "suffix"),
		field.Other("selected_value", &fieldjson.DeliveryRuleConstraintSelectedValue{}).
			SchemaType(map[string]string{
				dialect.SQLite:   "json",
				dialect.Postgres: "jsonb", // B otherwise we distinct causes the query to fail
			}),
	}
}

func (DeliveryRuleConstraint) Hooks() []ent.Hook {
	return []ent.Hook{
		constraint.TrimDeliveryRuleConstraint(),
		constraint.SaveDeliveryOptionConstraintCreate(),
		constraint.SaveDeliveryOptionConstraintUpdate(),
	}
}

func (DeliveryRuleConstraint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery_rule_constraint_group", DeliveryRuleConstraintGroup.Type).
			Ref("delivery_rule_constraints").
			Unique().
			Required().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (DeliveryRuleConstraint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("delivery_rule_constraint")),
	}
}
