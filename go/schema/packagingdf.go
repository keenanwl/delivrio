package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PackagingDF struct {
	ent.Schema
}

func (PackagingDF) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (PackagingDF) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("packaging", Packaging.Type).
			Ref("packaging_df").
			Required().
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}

func (PackagingDF) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (PackagingDF) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("packaging_df")),
	}
}

func (PackagingDF) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("api_type").
			Values("PKK", "PL1", "PL2", "PL4", "K10", "K20", "C10", "PL7", "CLL", "PLL"),
		field.Float("max_weight").
			Optional(),
		field.Float("min_weight").
			Optional(),
		field.Bool("stackable").
			Comment("Some carriers this is not boolean, so this is not a general toggle").
			Default(false),
	}
}
