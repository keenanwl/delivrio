package schema

import (
	"delivrio.io/go/schema/delivrioannotations"
	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

func (Address) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		delivrioannotations.Check(),
		delivrioannotations.Clone(),
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("uniqueness_id").
			Unique().
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipAll),
				delivrioannotations.Skip(),
				delivrioannotations.SkipClone(),
			),
		field.String("first_name"),
		field.String("last_name"),
		field.String("email"),
		field.String("phone_number"),
		field.String("phone_number_2").
			Optional().
			Comment("Some applications have both mobile and generic"),
		field.String("vat_number").
			Comment("Electronic customs").
			Optional(),
		field.String("company").
			Optional(),
		field.String("address_one"),
		field.String("address_two"),
		field.String("city"),
		field.String("state").Optional(),
		field.String("zip"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipient_colli", Colli.Type).
			Ref("recipient"),
		edge.From("recipient_consolidation", Consolidation.Type).
			Ref("recipient").
			Unique(),
		edge.From("company_address", Tenant.Type).
			Ref("company_address"),
		edge.From("location", Location.Type).
			Ref("address"),
		edge.From("sender_colli", Colli.Type).
			Ref("sender"),
		edge.From("sender_consolidation", Consolidation.Type).
			Ref("sender").
			Unique(),
		edge.From("return_sender_colli", ReturnColli.Type).
			Ref("sender"),
		edge.From("return_recipient_colli", ReturnColli.Type).
			Ref("recipient"),
		edge.To("country", Country.Type).
			Required().
			Unique().
			Annotations(delivrioannotations.Check()),
	}
}

func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("address")),
	}
}
