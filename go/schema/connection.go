package schema

import (
	"delivrio.io/go/schema/hooks/connectionhooks"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Connection struct {
	ent.Schema
}

func (Connection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

type ErrConnectionFields struct {
	InvalidFields map[string]string
}

func (ErrConnectionFields) Error() string {
	return "one or more connection fields is invalid"
}
func (e ErrConnectionFields) String() string {
	return e.Error()
}

var (
	camel = gen.Funcs["camel"].(func(string) string)
)

func (Connection) Hooks() []ent.Hook {
	return []ent.Hook{
		connectionhooks.UpdateCurrencyDependencies(),
	}
}

// Fields of the Connection.
func (Connection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("sync_orders").
			Default(false),
		field.Bool("sync_products").
			Default(false),
		field.Bool("fulfill_automatically").
			Default(false),
		field.Bool("dispatch_automatically").
			Default(false),
		field.Bool("convert_currency").
			Default(false),
		field.Bool("auto_print_parcel_slip").
			Default(false),
	}
}

// Edges of the Connection.
func (Connection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connection_brand", ConnectionBrand.Type).
			Required().
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		edge.To("connection_shopify", ConnectionShopify.Type).
			Unique(),
		edge.To("orders", Order.Type),
		edge.To("sender_location", Location.Type).
			Unique().
			Required(),
		edge.To("pickup_location", Location.Type).
			Unique().
			Required(),
		edge.To("return_location", Location.Type).
			Unique().
			Required(),
		edge.To("seller_location", Location.Type).
			Unique().
			Required(),
		edge.To("delivery_option", DeliveryOption.Type),
		edge.To("default_delivery_option", DeliveryOption.Type).
			Comment("Delivery option to be set when none specified via sync or API").
			Unique(),
		edge.From("return_portal", ReturnPortal.Type).
			Unique().
			Ref("connection"),
		edge.From("hypothesis_test", HypothesisTest.Type).
			Ref("connection"),
		edge.From("notifications", Notification.Type).
			Ref("connection"),
		edge.To("currency", Currency.Type).
			Unique().
			Required(),
		edge.To("packing_slip_template", Document.Type).
			Unique(),
		edge.From("connection_lookup", ConnectionLookup.Type).
			Ref("connections"),
	}
}

func (Connection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connection")),
	}
}
