package schema

import (
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type PackagingUSPSProcessingCategory struct {
	ent.Schema
}

func (PackagingUSPSProcessingCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (PackagingUSPSProcessingCategory) Edges() []ent.Edge {
	return []ent.Edge{
		//
	}
}

func (PackagingUSPSProcessingCategory) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (PackagingUSPSProcessingCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("packaging_usps_processing_category")),
	}
}

func (PackagingUSPSProcessingCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("processing_category").
			Values("LETTERS", "FLATS", "MACHINABLE", "IRREGULAR", "NON_MACHINABLE"),
	}
}
