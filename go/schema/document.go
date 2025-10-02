package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Document struct {
	ent.Schema
}

func (Document) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carrier_brand", CarrierBrand.Type).
			Unique(),
		edge.From("connection_packing_slip", Connection.Type).
			Ref("packing_slip_template"),
	}
}

func (Document) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("document")),
	}
}

func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("html_template").
			Optional(),
		field.String("html_header").
			Optional(),
		field.String("html_footer").
			Optional(),
		field.String("last_base64_pdf").
			Optional().
			Comment("Facilitates printing by saving the latest version of this document"),
		field.Enum("merge_type").
			Values("Orders", "PackingSlip", "Waybill").
			Default("Orders"),
		field.Enum("paper_size").
			Values("A4", "Four_x_six").
			Default("A4"),
		field.Time("start_at").
			Default(func() time.Time {
				now := time.Now()
				return time.Date(now.Year(), now.Month(), now.Day()-1, 7, 0, 0, 0, time.Local)
			}),
		field.Time("end_at").
			Default(func() time.Time {
				now := time.Now()
				return time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.Local)
			}),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
	}
}
