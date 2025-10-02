package schema

import (
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Printer struct {
	ent.Schema
}

func (Printer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Printer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workstation", Workstation.Type).
			Ref("printer").
			Unique().
			Required(),
		edge.From("print_jobs", PrintJob.Type).
			Ref("printer"),
	}
}

func (Printer) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Printer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("printer")),
	}
}

func (Printer) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_id").
			Comment("ID from desktop print client").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			GoType(pulid.ID("")).
			Unique(),
		field.String("name"),
		field.Bool("label_zpl").
			Default(false),
		field.Bool("label_pdf").
			Default(false),
		field.Bool("label_png").
			Default(false),
		field.Bool("document").
			Default(false),
		field.Bool("rotate_180").
			Default(false),
		field.Bool("use_shell").
			Default(false),
		field.Enum("print_size").
			Values("A4", "cm_100_150", "cm_100_192").
			Default("A4"),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("last_ping").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}
