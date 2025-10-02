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

// SignupOptions holds the schema definition for the SignupOptions entity.
type SignupOptions struct {
	ent.Schema
}

func (SignupOptions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Fields of the SignupOptions.
func (SignupOptions) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("better_delivery_options"),
		field.Bool("improve_pick_pack"),
		field.Bool("shipping_label"),
		field.Bool("custom_docs"),
		field.Bool("reduced_costs"),
		field.Bool("easy_returns"),
		field.Bool("click_collect"),
		field.Int("num_shipments"),
	}
}

// Edges of the SignupOptions.
func (SignupOptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("signup_options").
			Unique().
			Required(),
	}
}

func (SignupOptions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("signup_options")),
	}
}
