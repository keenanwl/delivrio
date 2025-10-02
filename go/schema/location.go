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
	"entgo.io/ent/schema/index"
)

// Location exists just to handle all the FKs
type Location struct {
	ent.Schema
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		delivrioannotations.Clone(),
	}
}

func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).
			Unique().
			Required().Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		edge.To("location_tags", LocationTag.Type).
			Required(),
		edge.From("sender_connection", Connection.Type).
			Ref("sender_location"),
		edge.From("pickup_connection", Connection.Type).
			Ref("pickup_location"),
		edge.From("return_connection", Connection.Type).
			Ref("return_location"),
		edge.From("seller_connection", Connection.Type).
			Ref("seller_location"),
		edge.From("return_portal", ReturnPortal.Type).
			Ref("return_location"),
		edge.From("delivery_option", DeliveryOption.Type).
			Ref("click_collect_location"),
		edge.From("colli", Colli.Type).
			Ref("click_collect_location"),
	}
}

func (Location) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("location")),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "tenant_id").Unique(),
	}
}

func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("Location name, not used in the address"),
	}
}
