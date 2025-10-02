package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type PackagingUSPS struct {
	ent.Schema
}

func (PackagingUSPS) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (PackagingUSPS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("packaging", Packaging.Type).
			Ref("packaging_usps").
			Required().
			Unique().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("packaging_usps_rate_indicator", PackagingUSPSRateIndicator.Type).
			Unique().
			Required(),
		edge.To("packaging_usps_processing_category", PackagingUSPSProcessingCategory.Type).
			Unique().
			Required(),
	}
}

func (PackagingUSPS) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (PackagingUSPS) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("packaging_usps")),
	}
}

func (PackagingUSPS) Fields() []ent.Field {
	return []ent.Field{}
}
