package schema

import (
	"delivrio.io/go/ent/currency"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Currency holds the schema definition for the Currency entity.
type Currency struct {
	ent.Schema
}

func (Currency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// Fields of the Currency.
func (Currency) Fields() []ent.Field {
	return []ent.Field{
		field.String("display").
			Unique(),
		field.Enum("currency_code").
			Values("DKK", "EUR", "SEK", "USD").
			Default("DKK"),
	}
}

// Edges of the Currency.
func (Currency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_line", OrderLine.Type).
			Ref("currency"),
		edge.From("delivery_rule", DeliveryRule.Type).
			Ref("currency"),
	}
}

func (Currency) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(currency.FieldCurrencyCode).Unique(),
	}
}

func (Currency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("currency")),
	}
}
