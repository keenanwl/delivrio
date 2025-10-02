package schema

import (
	"delivrio.io/go/schema/mixins"
	"time"

	"delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Workstation struct {
	ent.Schema
}

func (Workstation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Workstation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("printer", Printer.Type),
		edge.To("user", User.Type).
			Unique().
			Comment("the user who created the workstation").
			Immutable(),
		edge.From("selected_user", User.Type).
			Unique().
			Ref("selected_workstation").
			Comment("the user currently sending print jobs to this workstation"),
	}
}

func (Workstation) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (Workstation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("workstation")),
		mixins.ArchiveMixin{},
	}
}

func (Workstation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("device_type").
			Values("label_station", "app").
			Default("label_station"),
		field.String("registration_code").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable().
			Sensitive(),
		field.String("workstation_id").
			GoType(pulid.ID("")).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			).
			Immutable(),
		field.Time("last_ping").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
		field.Enum("status").
			Values("pending", "active", "offline", "disabled").
			Default("pending"),
		field.Bool("auto_print_receiver").
			Default(false),
	}
}
