package schema

import (
	"context"
	"net/url"
	"time"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/schema/hooks/dlvshopify"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ConnectionShopify holds the schema definition for the ConnectionShopify entity.
type ConnectionShopify struct {
	ent.Schema
}

func (ConnectionShopify) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

type ErrConnectionShopifyFields struct {
	InvalidFields map[string]string
}

func (ErrConnectionShopifyFields) Error() string {
	return "one or more Shopify fields are invalid"
}
func (e ErrConnectionShopifyFields) String() string {
	return e.Error()
}
func NewErrConnectionShopifyFields() ErrConnectionShopifyFields {
	return ErrConnectionShopifyFields{
		InvalidFields: map[string]string{},
	}
}

func (ConnectionShopify) Hooks() []ent.Hook {

	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ConnectionShopifyFunc(func(ctx context.Context, m *ent2.ConnectionShopifyMutation) (ent.Value, error) {
					validationErrs := NewErrConnectionShopifyFields()
					if u, ok := m.StoreURL(); ok {
						_, err := url.Parse(u)
						if err != nil {
							validationErrs.InvalidFields[camel(connectionshopify.FieldStoreURL)] = "store URL is not a valid URL; expected: https://my-store-myshopify.com"
						}
					}

					if len(validationErrs.InvalidFields) == 0 {
						return next.Mutate(ctx, m)
					}
					return nil, validationErrs
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		dlvshopify.CreateShopifyConnection(),
		dlvshopify.UpdateShopifyConnection(),
	}

}

func (ConnectionShopify) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("rate_integration").
			Comment("Since certain plans don't allow for external rates").
			Default(false),
		field.String("store_url").
			Optional().
			Unique(),
		field.String("api_key").
			Optional(),
		field.String("lookup_key").
			Optional().
			Comment("Used for token-authenticating Shopify rate lookups").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.Time("sync_from").
			Optional().
			Default(time.Now),
		field.Strings("filter_tags").
			Default([]string{}).
			Optional().
			Comment("When set, only orders with these tags will be synchronized. Supports a "),
	}
}

// Edges of the ConnectionShopify.
func (ConnectionShopify) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connection", Connection.Type).
			Ref("connection_shopify").
			Unique().
			Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (ConnectionShopify) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("connection_shopify")),
	}
}
