package schema

import (
	"delivrio.io/go/ent/privacy"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("tenant")),
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("vat_number").Optional(),
		field.String("invoice_reference").Optional(),
	}
}

func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connect_option_carriers", ConnectOptionCarrier.Type),
		edge.To("connect_option_platforms", ConnectOptionPlatform.Type),
		edge.From("plan", Plan.Type).
			Ref("tenant").
			Unique().
			// Sets global default on create
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)).
			Required(),
		edge.To("company_address", Address.Type).
			Unique(),
		edge.To("default_language", Language.Type).
			Required().
			// Sets global default on create
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)).
			Unique(),
		edge.To("billing_contact", Contact.Type).
			Unique(),
		edge.To("admin_contact", Contact.Type).
			Unique(),
	}
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// For Tenant type, we only allow admin users to mutate
			// the tenant information and deny otherwise.
			//rule.AllowIfAdmin(),
			//privacy.AlwaysDenyRule(),
		},
	}
}
