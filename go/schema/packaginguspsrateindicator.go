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

type PackagingUSPSRateIndicator struct {
	ent.Schema
}

func (PackagingUSPSRateIndicator) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

func (PackagingUSPSRateIndicator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("packaging_usps", PackagingUSPS.Type).
			Ref("packaging_usps_rate_indicator"),
	}
}

func (PackagingUSPSRateIndicator) Hooks() []ent.Hook {
	return []ent.Hook{
		//
	}
}

func (PackagingUSPSRateIndicator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		pulid.MixinWithPrefix(pulid_server_prefix.TypeToPrefix("packaging_usps_rate_indicator")),
	}
}

func (PackagingUSPSRateIndicator) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
	}
}
