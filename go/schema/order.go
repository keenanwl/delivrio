package schema

import (
	"delivrio.io/go/schema/fieldjson"
	"time"

	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/schema/hooks/accessrights"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/schema/interceptors/limiter"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_public_id"),
		field.String("external_id").
			Optional().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		field.String("comment_internal").Optional(),
		field.String("comment_external").Optional(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("email_sync_confirmation_at").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Enum("status").
			Values(
				"Pending",
				"Partially_dispatched",
				"Dispatched",
				"Cancelled",
			).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.JSON("note_attributes", fieldjson.NoteAttributes{}).
			Optional().
			Default(fieldjson.NoteAttributes{}).
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("tenant").
			Fields("order_public_id").
			Unique(),
		index.Edges("tenant").
			Fields("external_id").
			Unique(),
		index.Fields(order.FieldCreatedAt),
	}
}

func (Order) Hooks() []ent.Hook {
	return []ent.Hook{
		history.OrderCreate(),
		history.OrderUpdate(),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("order")),
	}
}

func (Order) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		limiter.Limiter(15),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_history", OrderHistory.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.From("connection", Connection.Type).
			Ref("orders").
			Unique().
			Required(),
		edge.To("colli", Colli.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("return_colli", ReturnColli.Type),
		edge.From("hypothesis_test_delivery_option_request", HypothesisTestDeliveryOptionRequest.Type).
			Ref("order").
			Unique(),
		edge.From("pallet", Pallet.Type).
			Ref("orders").
			Comment("Orders may be added to consolidation either through a pallet or directly").
			Unique(),
		edge.From("consolidation", Consolidation.Type).
			Ref("orders").
			Comment("Orders may be added to consolidation either through a pallet or directly").
			Unique(),
	}
}

func (Order) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			accessrights.CheckAccessRightsMutation("orders"),
		},
		Query: privacy.QueryPolicy{
			accessrights.CheckAccessRightsQuery("orders"),
		},
	}
}
